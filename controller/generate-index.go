package controller

import (
	"context"
	"io/fs"
	"log"
	"math"
	"path/filepath"
	"time"

	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
	"golang.org/x/sync/semaphore"
)

const maxWorkers int64 = 1000
const retryWaitTime = 125 * time.Millisecond

// GenerateIndex triggers the index generation process
func GenerateIndex(dao *daos.Dao, owner string) error {
	indexRebuilding, err := dao.FindFirstRecordByData("systemstatus", "name", "index_rebuilding")
	if err != nil {
		return err
	}
	indexRebuilding.Set("value", "true")
	dao.SaveRecord(indexRebuilding)

	// Simulate index rebuilding
	err = traversDirAndBuildIndex(dao, "data", owner)
	if err != nil {
		return err
	}

	indexRebuilding.Set("value", "false")
	dao.SaveRecord(indexRebuilding)

	return nil
}

// updateRecord updates or creates a record in the database
func updateRecord(dao *daos.Dao, path string, info fs.FileInfo, owner string) error {
	collection, err := dao.FindCollectionByNameOrId("data_resources")
	if err != nil {
		return err
	}

	dbParent, _ := dao.FindFirstRecordByData(collection.Id, "path", path)

	record := models.NewRecord(collection)
	record.Set("resourceName", info.Name())
	record.Set("path", path)
	record.Set("editors", owner)
	record.Set("type", "file")

	if dbParent != nil {
		record.Set("parent", dbParent.Id)
	}

	if info.IsDir() {
		record.Set("type", "dir")
	}

	dao.SaveRecord((record))

	return nil
}

// retryWithExponentialBackoff retries action with exponential backoff
func retryWithExponentialBackoff(action func() error, maxRetries int, initialDelay time.Duration, backoffFactor float64) error {
	retryCount := 0
	for {
		err := action()
		if err == nil {
			return nil
		}

		retryCount++
		if retryCount > maxRetries {
			return err // Return the last error
		}

		delay := time.Duration(float64(initialDelay) * math.Pow(backoffFactor, float64(retryCount-1)))
		time.Sleep(delay)
	}
}

// traversDirAndBuildIndex walks through the directory and builds the index
func traversDirAndBuildIndex(dao *daos.Dao, path string, owner string) error {
	sem := semaphore.NewWeighted(maxWorkers)
	ctx := context.TODO()

	err := filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		if err := sem.Acquire(ctx, 1); err != nil {
			return err
		}

		go func() {
			defer sem.Release(1)

			err := retryWithExponentialBackoff(func() error {
				return updateRecord(dao, path, info, owner)
			}, 5, retryWaitTime, 1.25)

			if err != nil {
				log.Printf("Failed to update record after retries: %v", err)
			}
		}()

		return nil
	})

	if err != nil {
		return err
	}

	// Wait for all goroutines to finish
	for sem.Acquire(ctx, maxWorkers) != nil {
		time.Sleep(retryWaitTime)
	}

	return nil
}
