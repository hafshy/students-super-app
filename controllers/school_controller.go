package controllers

import (
	"fmt"

	"github.com/hafshy/students-super-app/models"
	"github.com/hafshy/students-super-app/repositories"
	"github.com/labstack/echo/v4"
)

func AddSchool(c echo.Context) error {
    // Parse request body
	schoolRequest := new(models.SchoolRequest)
	err := c.Bind(schoolRequest)

    if err != nil {
        return c.JSON(500, models.BaseResponse {
			Status: false,
            Message: err.Error(),
			Data: nil,
		})
    }

    // Create School
    school := &models.School{
        Name:    schoolRequest.Name,
        Address: schoolRequest.Address,
        City:    schoolRequest.City,
        Country: schoolRequest.Country,
    }
	fmt.Println(schoolRequest)
	// fmt.Println(school)

	err = repositories.AddSchool(school)

    // Save the school in the database
    if err != nil {
        return c.JSON(500, models.BaseResponse {
			Status: false,
            Message: err.Error(),
			Data: nil,
		})
    }

	return c.JSON(200, models.BaseResponse{
		Status:  true,
		Message: "Successfully created",
		Data:    school,
	})
}
