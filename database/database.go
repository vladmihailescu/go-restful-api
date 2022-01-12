package database

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/vladmihailescu/go-restful-api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

func InitDatabase() error {
	var err error

	DBConn, err = gorm.Open(sqlite.Open("schema.db"), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("unable to init db: %w", err)
	}
	log.Debugf("database initialised")

	DBConn.AutoMigrate(&models.User{})
	log.Debugf("database migrated")

	return err
}
