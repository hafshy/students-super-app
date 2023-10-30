package routes

import (
	"github.com/hafshy/students-super-app/controllers"
	"github.com/labstack/echo/v4"
)


func InitRoute(e *echo.Echo) {
	e.POST("/school", controllers.AddSchool)
	e.POST("/student", controllers.AddStudent)
}
