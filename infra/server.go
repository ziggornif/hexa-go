package infra

import (
	"fmt"

	"hexa-go/handlers"
	"hexa-go/infra/config"
	"hexa-go/infra/monitoring"
	"hexa-go/infra/storage"
	"os"

	"github.com/gin-contrib/cors"
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
	s.logger.Info("[server] Run - starting...")
	conf, err := config.LoadConfig(".", s.logger)
	if err != nil {
		s.exit(err)
	}

	conf.ValidateConfig()

	db, errDB := storage.DBConnect(conf.GetConfig(), s.logger)

	if errDB != nil {
		s.exit(errDB)
	}

	router := gin.Default()
	router.Use(cors.Default())

	apiRouter := router.Group("/api")
	err = handlers.MakeHandlers(apiRouter, db, s.logger)
	if err != nil {
		s.exit(err)
	}

	monitoring.PrometheusHandler(router)
	monitoring.HeartbeatHandler(router)

	err = router.Run(fmt.Sprintf(":%v", conf.GetConfig().Port))
	if err != nil {
		s.exit(err)
	}
}

func (s *server) exit(err error) {
	s.logger.Error("Could not start app ...", err)
	os.Exit(1)
}

// Shutdown - shutdown server
func (s *server) Shutdown() {
	storage.DBClose()
}
