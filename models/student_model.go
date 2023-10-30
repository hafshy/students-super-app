package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
    ID        int		`json:"id" gorm:"primaryKey"`
	Name      string	`json:"name" gorm:"not null"`
	Gender string	`json:"gender" gorm:"not null"`
	Email     string    `json:"email" gorm:"not null"`
	PhoneNumber string `json:"phone_number" gorm:"not null"`
	EnrollYear int	`json:"enroll_year" gorm:"not null"`
	ExpectedGraduatedYear int	`json:"expected_graduated_year"`
	CurrentSemester int	`json:"current_semester" gorm:"not null"`
	SchoolID int `json:"school_id" gorm:"not null"`
	StudentAccount StudentAccount `json:"student_account" gorm:"not null;foreignKey:StudentID"`
}
