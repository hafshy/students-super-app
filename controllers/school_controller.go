package controllers

import (
	"fmt"
	"strconv"

	"github.com/hafshy/students-super-app/models"
	"github.com/hafshy/students-super-app/repositories"
	"github.com/labstack/echo/v4"
)

func RegisterSchool(c echo.Context) error {
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

func GetSchools(c echo.Context) error {
	var schools []models.School
	err := repositories.GetSchools(&schools)
    if err!= nil {
        return c.JSON(500, models.BaseResponse{
            Status:  false,
            Message: "Failed get data",
            Data:    nil,
        })
    }

    return c.JSON(200, models.BaseResponse{
        Status:  true,
        Message: "Successfully get data",
        Data:    schools,
    })
}

func GetSchool(c echo.Context) error {
	schoolID, err := strconv.Atoi(c.Param("id"))
	
    if err!= nil {
        return c.JSON(500, models.BaseResponse{
            Status:  false,
            Message: "Failed parse data",
            Data:    nil,
        })
    }

	var school models.School
	err = repositories.GetSchool(schoolID, &school)

	if err!= nil {
        return c.JSON(500, models.BaseResponse{
            Status:  false,
            Message: "Failed get data",
            Data:    nil,
        })
    }

    return c.JSON(200, models.BaseResponse{
        Status:  true,
        Message: "Successfully get data",
        Data:    school,
    })
} 
