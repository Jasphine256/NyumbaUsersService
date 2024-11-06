package database

import (
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() error {
	dsn := "host=localhost user=postgres password=Mikhael@Jasphine2611 dbname=NyumbaUsersDb port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Use appropriate logging level
	})
	if err != nil {
		return err
	}

	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}

	// Configure the connection pool
	sqlDB.SetMaxOpenConns(25)           // Max open connections
	sqlDB.SetMaxIdleConns(10)           // Max idle connections
	sqlDB.SetConnMaxLifetime(time.Hour) // Max connection lifetime

	return nil
}

var (
	once sync.Once
	db   *gorm.DB
)

func GetDB() *gorm.DB {
	once.Do(func() {
		db = DB
	})
	return db
}
