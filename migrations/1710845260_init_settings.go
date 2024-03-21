package migrations

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models/schema"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db)

		settings, _ := dao.FindSettings()
		settings.Meta.AppName = "Yos"
		settings.Logs.MaxDays = 30

		err := dao.SaveSettings(settings)

		if err != nil {
			return err
		}

		users, err := dao.FindCollectionByNameOrId("users")
		if err != nil {
			return err
		}

		users.Schema.AddField(&schema.SchemaField{
			Type: schema.FieldTypeBool,
			Name: "isAdmin",
		})

		return dao.SaveCollection(users)

	}, nil)
}
