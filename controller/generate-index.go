package controller

import (
	"os"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/models"
)

func GenerateIndex(app *pocketbase.PocketBase, owner string) error {
	indexRebuilding, err := app.Dao().FindFirstRecordByData("systemstatus", "name", "index_rebuilding")
	if err != nil {
		return err
	}
	indexRebuilding.Set("value", "true")
	app.Dao().SaveRecord(indexRebuilding)

	// simulate index rebuilding
	traversDirAndBuildIndex(app, "data", owner)

	indexRebuilding.Set("value", "false")
	app.Dao().SaveRecord(indexRebuilding)

	return nil
}

func traversDirAndBuildIndex(app *pocketbase.PocketBase, path string, owner string) error {
	entries, _ := os.ReadDir(path)
	collection, err := app.Dao().FindCollectionByNameOrId("data_resources")
	if err != nil {
		return err
	}

	parent, _ := app.Dao().FindFirstRecordByData(collection.Id, "path", path)

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
			app.Dao().SaveRecord((record))
			traversDirAndBuildIndex(app, path+"/"+entry.Name(), owner)
		}
		app.Dao().SaveRecord((record))
	}

	return nil
}
