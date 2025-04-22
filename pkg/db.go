package pkg

import (
	"github.com/Skapar/simple-rest/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func PingDB(db *gorm.DB, logger Logger) error {
	sqlDB, err := db.DB()
	if err != nil {
		logger.Error("Failed to get database instance: ", err)
		return err
	}
	if err := sqlDB.Ping(); err != nil {
		logger.Error("Failed to ping database: ", err)
		return err
	}
	return nil
}

func ConnectDB(config *config.Config, logger Logger) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.PostgresAddr), &gorm.Config{})
	if err != nil {
		logger.Error("Failed to connect to database: ", err)
		return nil, err
	}
	return db, nil
}
