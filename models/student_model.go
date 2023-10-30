package models

type Student struct {
	// gorm.Model
	ID        int		`json:"id" gorm:"primaryKey"`
	Name      string	`json:"name" gorm:"not null"`
	Gender string	`json:"gender" gorm:"not null"`
	Email     string    `json:"email" gorm:"not null"`
	PhoneNumber string `json:"phone_number" gorm:"not null"`
	EnrollYear int	`json:"enroll_year" gorm:"not null"`
	ExpectedGraduatedYear int	`json:"expected_graduated_year"`
	CurrentSemester int	`json:"current_semester" gorm:"not null"`
	SchoolID            int    `json:"school_id" gorm:"not null"`
	School              School `json:"school" gorm:"foreignKey:SchoolID"`
	// StudentAccount StudentAccount `json:"student_account" gorm:"foreignKey:ID"`
	StudentAccount StudentAccount `json:"student_account" gorm:"embedded"`
}
