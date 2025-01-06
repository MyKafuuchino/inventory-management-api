package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"inventory-management/config"
	"log"
	"sync"
)

var (
	DB   *gorm.DB
	once sync.Once
)

func InitDatabase() {
	once.Do(func() {
		var err error
		dbConfig := config.GlobalDbConfig
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			dbConfig.DbUser,
			dbConfig.DbPassword,
			dbConfig.DbHost,
			dbConfig.DbPort,
			dbConfig.DbName,
		)
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err != nil {
			log.Fatalf("Failed to connect to database : %v", err)
		}

		log.Println("Successfully connected to database")
	})
}
