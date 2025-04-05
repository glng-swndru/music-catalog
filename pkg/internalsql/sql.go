package internalsql

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(DataSourceName string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(DataSourceName), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v\n", err)
		return nil, err
	}
	return db, nil
}
