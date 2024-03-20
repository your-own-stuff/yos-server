package controller

import (
	"time"

	"github.com/pocketbase/pocketbase"
)

func GenerateIndex(app *pocketbase.PocketBase) error {
	indexRebuilding, err := app.Dao().FindFirstRecordByData("systemstatus", "name", "index_rebuilding")
	if err != nil {
		return err
	}
	indexRebuilding.Set("value", "true")
	app.Dao().SaveRecord(indexRebuilding)

	// simulate index rebuilding
	time.Sleep(5 * time.Second)

	indexRebuilding.Set("value", "false")
	app.Dao().SaveRecord(indexRebuilding)

	return nil
}
