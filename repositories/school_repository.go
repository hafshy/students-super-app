package repositories

import (
	"github.com/hafshy/students-super-app/configs"
	"github.com/hafshy/students-super-app/models"
)

func AddSchool(school *models.BaseResponse) error {
	result := configs.DB.Create(school)
	return result.Error
}
