package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(configOption ...ConfigOpt) (*gorm.DB, error) {
	var config Config
	for _, configure := range configOption {
		err := configure(&config)
		if err != nil {
			return nil, fmt.Errorf("fatal error configuring database: %v", err)
		}
	}

	dsn := config.DSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	return db, nil
}

func Must(configOption ...ConfigOpt) *gorm.DB {
	db, err := New(configOption...)
	if err != nil {
		panic(err)
	}
	return db
}
