package database

import (
	"fmt"
	"log"

	"github.com/oluwakeye-john/wallet-alert/config"
	"github.com/oluwakeye-john/wallet-alert/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupAndConnectDB() {
	dsn := config.MustGetEnv("DATABASE_URL")
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		log.Fatal("Failed to connect to database")
	}

	DB = conn
	fmt.Println("Connected to database")
}

func Migrate() {
	DB.AutoMigrate(
		&models.Address{},
		&models.Account{},
	)
}
