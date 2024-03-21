package controller

import (
	"os"

	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
)

func GenerateIndex(dao *daos.Dao, owner string) error {
	indexRebuilding, err := dao.FindFirstRecordByData("systemstatus", "name", "index_rebuilding")
	if err != nil {
		return err
	}
	indexRebuilding.Set("value", "true")
	dao.SaveRecord(indexRebuilding)

	// simulate index rebuilding
	traversDirAndBuildIndex(dao, "data", owner)

	indexRebuilding.Set("value", "false")
	dao.SaveRecord(indexRebuilding)

	return nil
}

func traversDirAndBuildIndex(dao *daos.Dao, path string, owner string) error {
	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	collection, err := dao.FindCollectionByNameOrId("data_resources")
	if err != nil {
		return err
	}

	parent, _ := dao.FindFirstRecordByData(collection.Id, "path", path)

	for _, entry := range entries {
		record := models.NewRecord(collection)
		record.Set("resourceName", entry.Name())
		record.Set("path", path+"/"+entry.Name())
		record.Set("editors", owner)
		record.Set("type", "file")
		if parent != nil {
			record.Set("parent", parent.Id)
		}
		if entry.IsDir() {
			record.Set("type", "dir")
			dao.SaveRecord((record))
			traversDirAndBuildIndex(dao, path+"/"+entry.Name(), owner)
		}
		dao.SaveRecord((record))
	}

	return nil
}
