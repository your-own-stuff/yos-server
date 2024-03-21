package migrations

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db)

		admin := &models.Admin{}
		admin.Email = "admin@admin.com"
		admin.SetPassword("adminadmin")

		err := dao.SaveAdmin(admin)

		if err != nil {
			return err
		}

		collection, err := dao.FindCollectionByNameOrId("users")
		if err != nil {
			return err
		}

		record := models.NewRecord(collection)
		record.SetUsername("admin")
		record.SetPassword("asdfasdf")
		record.Set("name", "Admin")
		record.Set("email", "user@probs.at")
		record.Set("isAdmin", true)
		record.SetVerified(true)

		return dao.SaveRecord(record)
	}, nil)
}
