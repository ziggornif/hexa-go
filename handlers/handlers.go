package handlers

import (
	"hexa-go/handlers/rest"
	"hexa-go/packages/todo"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
)

func makeRestHandler(route *gin.RouterGroup, db *pg.DB, logger *logrus.Logger) error {
	repo, err := todo.NewRepository(db, logger)
	if err != nil {
		return err
	}

	service := todo.NewService(repo, logger)
	rest.NewTodoController(route, service, logger)

	return nil
}

// MakeHandlers - create handlers
func MakeHandlers(route *gin.RouterGroup, db *pg.DB, logger *logrus.Logger) error {
	err := makeRestHandler(route, db, logger)
	return err
}
