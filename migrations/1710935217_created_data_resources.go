package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		jsonData := `{
			"id": "fvkd7jhc21136qr",
			"created": "2024-03-20 11:46:57.711Z",
			"updated": "2024-03-20 11:46:57.711Z",
			"name": "data_resources",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "x3t954g1",
					"name": "resourceName",
					"type": "text",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"min": null,
						"max": null,
						"pattern": ""
					}
				},
				{
					"system": false,
					"id": "98cefp2k",
					"name": "path",
					"type": "text",
					"required": true,
					"presentable": false,
					"unique": true,
					"options": {
						"min": null,
						"max": null,
						"pattern": ""
					}
				},
				{
					"system": false,
					"id": "auzuknwz",
					"name": "type",
					"type": "text",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"min": null,
						"max": null,
						"pattern": ""
					}
				},
				{
					"system": false,
					"id": "pskpudmz",
					"name": "reader",
					"type": "relation",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"collectionId": "_pb_users_auth_",
						"cascadeDelete": false,
						"minSelect": null,
						"maxSelect": null,
						"displayFields": null
					}
				},
				{
					"system": false,
					"id": "qvb1nnzq",
					"name": "editors",
					"type": "relation",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"collectionId": "_pb_users_auth_",
						"cascadeDelete": false,
						"minSelect": 1,
						"maxSelect": null,
						"displayFields": null
					}
				}
			],
			"indexes": [
				"CREATE UNIQUE INDEX ` + "`" + `idx_gwETy7z` + "`" + ` ON ` + "`" + `data_resources` + "`" + ` (` + "`" + `path` + "`" + `)"
			],
			"listRule": "@request.auth.id != \"\"",
			"viewRule": "@request.auth.id != \"\"",
			"createRule": null,
			"updateRule": null,
			"deleteRule": null,
			"options": {}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("fvkd7jhc21136qr")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
