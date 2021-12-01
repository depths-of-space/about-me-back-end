package db

import (
	"about-me-back-end/types"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Function for connect DB and generate Entites
func ConnectDB() gorm.DB {

	dsn := "host=127.0.0.1 user=about password=about_password dbname=about port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()

	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxIdleConns(2)

	sqlDB.SetMaxOpenConns(5)

	db.AutoMigrate(&types.Post{})
	db.AutoMigrate(&types.Tags{})

	tag := types.Tags{
		Subject: "javascript",
	}

	// Create
	db.Create(&types.Post{
		Title: "Meu primeiro post",
		Body:  "Esse post é o primeiro de muitos da minha rede social about-me",
		Tags: []types.Tags{
			tag,
		},
	})

	return *db
}
