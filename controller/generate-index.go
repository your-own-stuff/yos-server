package controller

import (
	"io/fs"
	"path/filepath"

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

func traversDirAndBuildIndex(dao *daos.Dao, path string, owner string) error {
	// number is basically a RAM vs Speed toggle
	// magic break even is when you have one routine per file
	var sem = make(chan string, 5000)

	err := filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		sem <- path
		go func() {
			updateRecord(dao, path, info, owner)
			<-sem
		}()

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
