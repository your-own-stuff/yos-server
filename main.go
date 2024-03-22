package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"yos/controller"

	_ "yos/migrations"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
)

func main() {
	app := pocketbase.New()

	// loosely check if it was executed using "go run"
	isGoRun := strings.HasPrefix(os.Args[0], os.TempDir())

	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		// enable auto creation of migration files when making collection changes in the Admin UI
		// (the isGoRun check is to enable it only during development)
		Automigrate: isGoRun,
	})

	// serves static files from the provided public dir (if exists)
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./pb_public"), false))
		e.Router.GET("/rebuild-index", func(c echo.Context) error {
			authRecord := apis.RequestInfo(c).AuthRecord
			// unauthorized
			if authRecord == nil || !authRecord.GetBool("isAdmin") {
				return c.JSON(http.StatusUnauthorized, "Unauthorized")
			}

			go controller.GenerateIndex(app.Dao(), authRecord.Id)

			return c.JSON(http.StatusOK, map[string]string{"status": "started"})
		})
		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
