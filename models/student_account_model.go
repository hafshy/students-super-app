package models


type StudentAccount struct {
	Username string	`json:"username" gorm:"not null"`
	Password string	`json:"password" gorm:"not null"`
}
