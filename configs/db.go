package configs

import (
	"fmt"
	"os"

	"github.com/hafshy/students-super-app/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", 
		os.Getenv("DB_USER"), 
		os.Getenv("DB_PASSWORD"), 
		os.Getenv("DB_HOST"), 
		os.Getenv("DB_PORT"), 
        os.Getenv("DB_NAME"),
	)
	fmt.Println(dsn)
	var err error
  	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err!= nil {
        panic("failed to connect database")
    }
	AutoMigrate()
}

func AutoMigrate() {
	DB.AutoMigrate(&models.Student{}, &models.School{})
}
