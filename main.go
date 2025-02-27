package main

import (
	"fmt"
	"online-shop/entity"
	"online-shop/helper"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	connectURI := os.Getenv("DB_CONNECTION_URI")

	db, err := gorm.Open(postgres.Open(connectURI), &gorm.Config{})
	if err != nil {
		helper.PanicIfError(err)
	}

	sqlDB, err := db.DB()

	if err != nil {
		helper.PanicIfError(err)
	}

	defer sqlDB.Close()

	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxIdleTime(15 * time.Minute)
	sqlDB.SetConnMaxLifetime(1 * time.Hour)

	db.AutoMigrate(&entity.Product{})

	product := entity.Product{
		Name:     "Keras A4",
		Category: "Kertas",
		Price:    1000000,
	}

	result := db.Create(&product)

	if result.Error != nil {
		helper.PanicIfError(err)
	}

	fmt.Println("Data berhasil di buat !")
}
