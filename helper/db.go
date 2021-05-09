package helper

import (
	"bwa-startup/campaign"
	"bwa-startup/transaction"
	"bwa-startup/user"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDb() *gorm.DB {
	dsn := os.Getenv("DSN_CONNECTION")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&user.User{}, &campaign.Campaign{}, &campaign.CampaignImage{}, &transaction.Transaction{})
	return db
}
