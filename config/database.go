package config

import (
	"fmt"
	"go-todo/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	// dsn stands for data source name
	dsn := "postgresql://postgres:Engkar.K@12@localhost:5432/db_go_todo?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("connection opened to database")

	return db
}

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.Item{},
		&models.SubItem{},
	)
	if err != nil {
		panic("failed to migrate database")
	}
	fmt.Println("database migrated")
}
