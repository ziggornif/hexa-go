package storage

import (
	"errors"
	"hexa-go/infra/config"

	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
)

var db *pg.DB

// DBConnect - create db connection
func DBConnect(config *config.Configuration, logger *logrus.Logger) (*pg.DB, error) {
	if db != nil {
		return db, nil
	}

	opts := &pg.Options{
		Addr:     config.DBURL,
		User:     config.DBUser,
		Password: config.DBPass,
		Database: config.DBName,
	}

	db = pg.Connect(opts)
	if db == nil {
		logger.Error("Fail to connect to database.")
		return nil, errors.New("fail to connect to database")
	}

	logger.Info("Connection to database successful.")
	return db, nil
}

// DBClose - close db connection
func DBClose() {
	_ = db.Close()
	db = nil
}
