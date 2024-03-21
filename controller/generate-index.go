package controller

import (
	"io/fs"
	"log"
	"math"
	"path/filepath"
	"time"

	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
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

type workItem struct {
	path string
	info fs.DirEntry
}

func retryWithExponentialBackoff(action func() error, maxRetries int, initialDelay time.Duration, backoffFactor float64) error {
	retryCount := 0
	for {
		err := action()
		if err == nil {
			return nil
		}

		retryCount++
		if retryCount > maxRetries {
			return err
		}

		delay := time.Duration(float64(initialDelay) * math.Pow(backoffFactor, float64(retryCount-1)))
		time.Sleep(delay)
	}
}

func worker(dao *daos.Dao, owner string, workQueue <-chan workItem, done chan<- bool) {
	for item := range workQueue {
		err := retryWithExponentialBackoff(func() error {
			return updateRecord(dao, item.path, item.info, owner)
		}, 5, retryWaitTime, 1.25)

		if err != nil {
			log.Printf("Failed to update record for %s: %v", item.path, err)
		}
	}
	done <- true
}

func updateRecord(dao *daos.Dao, path string, info fs.DirEntry, owner string) error {
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

func traversDirAndBuildIndex(dao *daos.Dao, path string, owner string) error {
	// Create work queue and done channel
	workQueue := make(chan workItem, maxWorkers) // Buffer based on maxWorkers
	done := make(chan bool)

	// Start a fixed number of worker goroutines
	for i := 0; i < int(maxWorkers); i++ {
		go worker(dao, owner, workQueue, done)
	}

	// Dispatch work to the work queue
	err := filepath.WalkDir(path, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		workQueue <- workItem{path: path, info: info}

		return nil
	})

	if err != nil {
		return err
	}

	// Close the work queue and wait for all workers to finish
	close(workQueue)

	for i := 0; i < int(maxWorkers); i++ {
		<-done // Wait for each worker to signal completion
	}

	return nil
}
