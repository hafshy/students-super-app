package repositories

import (
	"github.com/hafshy/students-super-app/configs"
	"github.com/hafshy/students-super-app/models"
)

func AddStudent(student *models.Student) error {
	result := configs.DB.Create(student)
	return result.Error
}

func GetStudent(studentID int, student *models.Student) error {
	result := configs.DB.First(&student, studentID)
    return result.Error
}

func GetStudents(students *[]models.Student) error {
	result := configs.DB.Find(&students)
    return result.Error
}

func AuthenticateStudent(account models.StudentAccount, student *models.Student) error {
	result := configs.DB.Where("username= ? AND password= ?", account.Username, account.Password).First(&student)
	return result.Error
}
