package models


type StudentAccount struct {
	// ID        int	`json:"id" gorm:"primaryKey"`
	Username string	`json:"username" gorm:"not null"`
	Password string	`json:"password" gorm:"not null"`
	// StudentID	int	`json:"student_id" gorm:"unique;not null;autoIncrement"`
}
