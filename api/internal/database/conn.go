package database

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SQLiteConn() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(os.Getenv("SQLITE_DB")), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		return nil, err
	}

	return db, nil
}
