package controllers

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/hafshy/students-super-app/models"
	"github.com/hafshy/students-super-app/repositories"
	"github.com/labstack/echo/v4"
)

func AddStudent(c echo.Context) error {
	req := new(models.RegisterRequest)

	if err := c.Bind(req); err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }

	var school models.School
	err := repositories.GetSchool(req.SchoolID, &school)
	if err != nil {
		return c.JSON(500, models.BaseResponse {
			Status: false,
            Message: "School not found",
			Data: nil,
		})
	}

	if !repositories.SchoolExists(req.SchoolID) {
        return c.JSON(500, models.BaseResponse {
			Status: false,
            Message: "School not found",
			Data: nil,
		})
    }

	student := &models.Student {
        Name:                 req.Name,
        Gender:               req.Gender,
        Email:                req.Email,
        PhoneNumber:          req.PhoneNumber,
        EnrollYear:           req.EnrollYear,
        ExpectedGraduatedYear: req.ExpectedGraduatedYear,
        CurrentSemester:      req.CurrentSemester,
		School: school,
        // SchoolID:             req.SchoolID,
        StudentAccount:       models.StudentAccount{
			Username: req.Username,
			Password: req.Password,
		},
    }

	err = repositories.AddStudent(student)
	if err != nil {
		return c.JSON(500, models.BaseResponse{
			Status:  false,
			Message: "Failed create student data",
			Data:    nil,
		})
	}

	// studentAccount := &models.StudentAccount{
    //     Username: req.Username,
    //     Password: req.Password,
    // }

	// err = repositories.AddStudentAccount(studentAccount)

	// if err != nil {
    //     return c.JSON(500, models.BaseResponse{
	// 		Status:  false,
	// 		Message: "Failed get create student account",
	// 		Data:    nil,
	// 	})
    // }


	token := jwt.New(jwt.SigningMethodHS256)
	claims := models.JwtClaims {
		Email: student.Email,
		Username: student.StudentAccount.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}
	token.Claims = claims

	signedToken, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.JSON(500, models.BaseResponse{
			Status:  false,
			Message: "Failed to create token",
			Data:    nil,
		})
	}

	return c.JSON(200, models.BaseResponse{
		Status:  true,
		Message: "Successfully get data",
		Data:    echo.Map {
			"token": signedToken,
		},
	})
}


