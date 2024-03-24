package yos

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/daos"
)

type Server struct {
	logger *slog.Logger
	dao    *daos.Dao
}

func New(logger *slog.Logger, dao *daos.Dao) *Server {
	if logger == nil {
		logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
	}

	return &Server{
		logger: logger,
		dao:    dao,
	}
}

type YosRoutes struct {
	Path        string
	Method      string
	HandlerFunc echo.HandlerFunc
}

func (y *Server) GetRoutes() []YosRoutes {
	return []YosRoutes{
		{Path: "/rebuild-index", Method: http.MethodGet, HandlerFunc: func(c echo.Context) error {
			authRecord := apis.RequestInfo(c).AuthRecord
			// unauthorized
			if authRecord == nil || !authRecord.GetBool("isAdmin") {
				return c.JSON(http.StatusUnauthorized, "Unauthorized")
			}

			go GenerateIndex(y.dao, authRecord.Id)

			return c.JSON(http.StatusOK, map[string]string{"status": "started"})
		}},
	}
}
