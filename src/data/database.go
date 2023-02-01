package data

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewDB - Creates a new database connection
func NewDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
