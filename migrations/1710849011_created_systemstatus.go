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
			"id": "zg8q4g3gl7xk0gd",
			"created": "2024-03-19 11:50:11.539Z",
			"updated": "2024-03-19 11:50:11.539Z",
			"name": "systemstatus",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "w02qgcif",
					"name": "name",
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
					"id": "fnq1b8fu",
					"name": "value",
					"type": "text",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"min": null,
						"max": null,
						"pattern": ""
					}
				}
			],
			"indexes": [
				"CREATE UNIQUE INDEX ` + "`" + `idx_iteV39f` + "`" + ` ON ` + "`" + `systemstatus` + "`" + ` (` + "`" + `name` + "`" + `)"
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

		err := daos.New(db).SaveCollection(collection)
		if err != nil {
			return err
		}

		dao, err := daos.New(db).FindCollectionByNameOrId("systemstatus")
		if err != nil {
			return err
		}

		record := models.NewRecord(dao)
		record.Set("name", "index_rebuilding")
		record.Set("value", "false")

		return daos.New(db).SaveRecord(record)

	}, nil)
}
