package db

import (
	"about-me-back-end/types"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectDB() {

	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	// Migrate the schema
	db.AutoMigrate(&types.Post{})
	db.AutoMigrate(&types.Tags{})

	// Create
	db.Create(&types.Post{})
	db.Create(&types.Tags{})

}
