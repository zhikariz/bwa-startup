package helper

import (
	"bwa-startup/user"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDb() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/bwa-startup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&user.User{})
	return db
}
