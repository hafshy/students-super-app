package models

type RegisterRequest struct {
	Name                string `json:"name"`
	Gender              string `json:"gender"`
	Email               string `json:"email"`
	PhoneNumber         string `json:"phone_number"`
	EnrollYear          int    `json:"enroll_year"`
	ExpectedGraduatedYear int    `json:"expected_graduated_year"`
	CurrentSemester      int    `json:"current_semester"`
	SchoolID            int    `json:"school_id"`
	Username            string `json:"username"`
	Password            string `json:"password"`
}
