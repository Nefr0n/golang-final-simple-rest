package pkg

import (
	"github.com/Skapar/simple-rest/config"
	"gorm.io/gorm"
)

type Database interface {
	Connect(config *config.Config, logger Logger) (*gorm.DB, error)
	Ping(db *gorm.DB, logger Logger) error
}

type GormDatabase struct{}

func (g *GormDatabase) Connect(config *config.Config, logger Logger) (*gorm.DB, error) {
	return ConnectDB(config, logger)
}

func (g *GormDatabase) Ping(db *gorm.DB, logger Logger) error {
	return PingDB(db, logger)
}
