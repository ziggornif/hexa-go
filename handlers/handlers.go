package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
)

// MakeHandlers - create handlers
func MakeHandlers(route *gin.RouterGroup, db *pg.DB, logger *logrus.Logger) error {
	return nil
}
