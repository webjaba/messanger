package storage

import (
	"db-service/internal/config"
	postgresql "db-service/internal/storage/postgre_sql"

	"gorm.io/gorm"
)

func ConnectDB(cfg *config.Config) *gorm.DB {
	switch cfg.DBEngine {
	case "postgre_sql":
		return postgresql.ConnectDB()
	}
	return nil
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Message{})
}
