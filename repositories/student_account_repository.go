package repositories

import (
	"github.com/hafshy/students-super-app/configs"
	"github.com/hafshy/students-super-app/models"
)

func AddStudentAccount(studentAccount *models.StudentAccount) error {
	result := configs.DB.Create(studentAccount)
	return result.Error
}
