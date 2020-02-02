package database

import (
	"github.com/jinzhu/gorm"
)

func Open() (db *gorm.DB) {
	return connection()
}

func Close() {
	db.Close()
}
