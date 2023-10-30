package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/hafshy/students-super-app/models"
	"github.com/hafshy/students-super-app/repositories"
	"github.com/labstack/echo/v4"
)

func RegisterStudent(c echo.Context) error {
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
		Message: "Successfully register",
		Data:    echo.Map {
			"token": signedToken,
		},
	})
}

func LoginStudent(c echo.Context) error {
	account := new(models.StudentAccount)
	err := c.Bind(account)

	if err!= nil {
        return c.JSON(500, models.BaseResponse {
			Status: false,
            Message: err.Error(),
			Data: nil,
		})
    }

	var student models.Student
	err = repositories.AuthenticateStudent(*account, &student)
	if err!= nil {
        return c.JSON(500, models.BaseResponse {
			Status: false,
            Message: "Invalid username or password",
			Data: nil,
		})
    }

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
		Message: "Successfully login",
		Data:    echo.Map {
			"token": signedToken,
		},
	})
}

func GetStudents(c echo.Context) error {
	var students []models.Student
	err := repositories.GetStudents(&students)
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
        Data:    students,
    })
}

func GetStudent(c echo.Context) error {
	studentID, err := strconv.Atoi(c.Param("id"))
	
    if err!= nil {
        return c.JSON(500, models.BaseResponse{
            Status:  false,
            Message: "Failed parse data",
            Data:    nil,
        })
    }

	var student models.Student
	err = repositories.GetStudent(studentID, &student)

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
        Data:    student,
    })
}
