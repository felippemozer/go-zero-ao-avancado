package database

import (
	"emailn/internal/domain/campaign"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDb() *gorm.DB {
	dsn := os.Getenv("DATABASE")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("fail to connect to database")
	}

	err = db.AutoMigrate(&campaign.Campaign{}, &campaign.Contact{})
	if err != nil {
		panic("error on initial database migration")
	} else {
		fmt.Println("Migration success")
	}

	return db
}
