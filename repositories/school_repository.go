package repositories

import (
	"github.com/hafshy/students-super-app/configs"
	"github.com/hafshy/students-super-app/models"
)

func AddSchool(school *models.School) error {
	result := configs.DB.Create(school)
	return result.Error
}

func GetSchool(schoolID int, school *models.School) error {
	result := configs.DB.First(&school, schoolID)
	return result.Error
}

func GetSchools(schools *[]models.School) error {
	result := configs.DB.Find(&schools)
    return result.Error
}

func SchoolExists(schoolID int) bool {
    var count int64
    configs.DB.Model(&models.School{}).Where("id = ?", schoolID).Count(&count)
    return count > 0
}
