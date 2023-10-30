package configs

import (
	"github.com/hafshy/students-super-app/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:@tcp(127.0.0.1:3306)/students_super_app?charset=utf8mb4&parseTime=True&loc=Local"
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
