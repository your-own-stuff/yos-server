package main

import (
	"log"
	"net/http"
	"os"
	"strings"
	"yos/types"

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
            admin := apis.RequestInfo(c).Admin

            if admin == nil {
                return c.JSON(http.StatusUnauthorized, "Unauthorized")
            }

            status := types.Systemstatus{}
            err := app.Dao().DB().NewQuery("SELECT * FROM systemstatus WHERE name = 'index_rebuilding'").One(&status)

            if err != nil {
                return c.JSON(http.StatusInternalServerError, err.Error())
            }

            if status.Value == "true" {
                return c.JSON(http.StatusConflict, "Index is already being rebuilt")
            }
            
            status.Value = "true"
            err = app.Dao().DB().Model(&status).Update("Value")

            if err != nil {
                log.Println(err)
                return c.JSON(http.StatusInternalServerError, err.Error())
            }
            return c.JSON(http.StatusOK, map[string]bool{"success": true})
        })
        return nil
    })

    if err := app.Start(); err != nil {
        log.Fatal(err)
    }
}