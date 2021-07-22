package infra

import (
	"hexa-go/handlers"
	"hexa-go/infra/config"
	"hexa-go/infra/monitoring"
	"hexa-go/infra/storage"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Server - server interface
type Server interface {
	Run()
	Shutdown()
}

type server struct {
	logger *logrus.Logger
}

// NewServer - create server instance
func NewServer(logger *logrus.Logger) Server {
	return &server{logger: logger}
}

// Run - run server
func (s *server) Run() {
	conf, err := config.LoadConfig(".", s.logger)
	if err != nil {
		log.Fatal("Unable to read the config file: ", err)
		return
	}

	conf.ValidateConfig()

	db, errDB := storage.DBConnect(conf.GetConfig(), s.logger)

	if errDB != nil {
		s.exit()
	}

	router := gin.Default()
	router.Use(cors.Default())
	router.Use(static.Serve("/", static.LocalFile("./static", false)))
	router.Use(static.Serve("/assets", static.LocalFile("./doc", false)))

	router.LoadHTMLFiles("./doc/swagger.html")
	router.GET("/docs", func(c *gin.Context) {
		c.HTML(http.StatusOK, "swagger.html", nil)
	})

	apiRouter := router.Group("/api")
	err = handlers.MakeHandlers(apiRouter, db, s.logger)
	if err != nil {
		s.exit()
	}

	monitoring.PrometheusHandler(router)
	monitoring.HeartbeatHandler(router)

	err = router.Run()
	if err != nil {
		s.exit()
	}
}

func (s *server) exit() {
	s.logger.Error("Could not start app ...") //TODO: improve this log message
	os.Exit(1)
}

// Shutdown - shutdown server
func (s *server) Shutdown() {
	storage.DBClose()
}
