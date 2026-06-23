package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConectarSQLite() (*gorm.DB, error) {

	db, err := gorm.Open(
		sqlite.Open("asociacion.db"),
		&gorm.Config{},
	)

	if err != nil {
		return nil, err
	}

	return db, nil
}
