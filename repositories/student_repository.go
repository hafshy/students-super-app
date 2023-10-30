package repositories

import (
	"github.com/hafshy/students-super-app/configs"
	"github.com/hafshy/students-super-app/models"
)

func AddStudent(student *models.Student) error {
	result := configs.DB.Create(student)
	return result.Error
}
