package migrations

import (
	"log"

	"github.com/caarlos0/env/v10"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

type InitAdminConfig struct {
	PBAdminEmail string `env:"PB_ADMIN_EMAIL,required"`
	PBAdminPW    string `env:"PB_ADMIN_PW,required"`

	YOSUsername string `env:"YOS_USERNAME,required"`
	YOSPassword string `env:"YOS_PASSWORD,required"`
	YOSName     string `env:"YOS_NAME,required"`
	YOSEmail    string `env:"YOS_EMAIL,required"`
}

func init() {

	m.Register(func(db dbx.Builder) error {
		cfg := InitAdminConfig{}
		if err := env.Parse(&cfg); err != nil {
			log.Fatal(err)
		}

		dao := daos.New(db)

		admin := &models.Admin{}
		admin.Email = cfg.PBAdminEmail
		admin.SetPassword(cfg.PBAdminPW)

		err := dao.SaveAdmin(admin)

		if err != nil {
			return err
		}

		collection, err := dao.FindCollectionByNameOrId("users")
		if err != nil {
			return err
		}

		record := models.NewRecord(collection)
		record.SetUsername(cfg.YOSUsername)
		record.SetPassword(cfg.YOSPassword)
		record.Set("name", cfg.YOSName)
		record.Set("email", cfg.YOSEmail)
		record.Set("isAdmin", true)
		record.SetVerified(true)

		return dao.SaveRecord(record)
	}, nil)
}
