package models


type School struct {
	// gorm.Model
	ID int `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"not null"`
	Address string `json:"address" gorm:"not null"`
	City string `json:"city" gorm:"not null"`
	Country string `json:"country" gorm:"not null"`
	// Students []Student `gorm:"foreignKey:SchoolID" json:"students"`
}
