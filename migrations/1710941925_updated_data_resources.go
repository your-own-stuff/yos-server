package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models/schema"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("fvkd7jhc21136qr")
		if err != nil {
			return err
		}

		// add
		newField := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "rq9vuzbo",
			"name": "parent",
			"type": "relation",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "fvkd7jhc21136qr",
				"cascadeDelete": true,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": null
			}
		}`), newField); err != nil {
			return err
		}
		collection.Schema.AddField(newField)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("fvkd7jhc21136qr")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("rq9vuzbo")

		return dao.SaveCollection(collection)
	})
}
