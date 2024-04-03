package yos

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/daos"
)

const maxWorkers int64 = 1000
const retryWaitTime = 125 * time.Millisecond

type Server struct {
	logger    *slog.Logger
	dao       *daos.Dao
	workQueue chan WorkItem
	done      chan bool
}

func New(logger *slog.Logger) *Server {
	if logger == nil {
		logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
	}

	return &Server{
		logger:    logger,
		dao:       nil,
		workQueue: make(chan WorkItem, maxWorkers),
		done:      make(chan bool),
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

			go GenerateIndex(y, authRecord.Id)

			return c.JSON(http.StatusOK, map[string]string{"status": "started"})
		}},
	}
}

func (y *Server) SetDao(dao *daos.Dao) {
	y.dao = dao
}

func (y *Server) Start() {
	// go Worker(y.dao, "system", y.workQueue, y.doneSig)
	// Start a fixed number of worker goroutines
	for i := 0; i < int(maxWorkers); i++ {
		go Worker(y.workQueue, y.done)
	}
}

func (y *Server) Stop() {
	// Close the work queue and wait for all workers to finish
	close(y.workQueue)

	for i := 0; i < int(maxWorkers); i++ {
		<-y.done // Wait for each worker to signal completion
	}
}
