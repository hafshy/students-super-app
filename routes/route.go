package routes

import (
	"github.com/hafshy/students-super-app/controllers"
	"github.com/labstack/echo/v4"
)


func InitRoute(e *echo.Echo) {
	e.POST("/school", controllers.RegisterSchool)
	e.POST("/student", controllers.RegisterStudent)
	e.POST("/login", controllers.LoginStudent)
	e.GET("/students", controllers.GetStudents)
	e.GET("/students/:id", controllers.GetStudent)
	e.GET("/schools", controllers.GetSchools)
	e.GET("/schools/:id", controllers.GetSchool)
}
