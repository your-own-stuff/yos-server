package yos

import (
	"io/fs"
	"path/filepath"

	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
)

// GenerateIndex triggers the index generation process
func GenerateIndex(yos *Server, owner string) error {
	indexRebuilding, err := yos.dao.FindFirstRecordByData("systemstatus", "name", "index_rebuilding")

	if err != nil {
		return err
	}

	indexRebuilding.Set("value", "true")
	yos.dao.SaveRecord(indexRebuilding)

	// Simulate index rebuilding
	err = traversDirAndBuildIndex(yos, "data", owner)

	if err != nil {
		return err
	}

	indexRebuilding.Set("value", "false")
	yos.dao.SaveRecord(indexRebuilding)

	return nil
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

func traversDirAndBuildIndex(yos *Server, path string, owner string) error {

	// Dispatch work to the work queue
	err := filepath.WalkDir(path, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// workQueue <- workItem{path: path, info: info}
		yos.workQueue <- WorkItem{
			action: func() error {
				return updateRecord(yos.dao, path, info, owner)
			},
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
