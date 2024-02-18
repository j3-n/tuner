package database

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SQLiteConn() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(os.Getenv("SQLITE_DB")), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
