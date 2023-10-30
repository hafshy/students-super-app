package main

import (
	"fmt"
	"os"

	"github.com/hafshy/students-super-app/configs"
	"github.com/hafshy/students-super-app/routes"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	loadEnv()
	configs.InitDB()
	e := echo.New()
	routes.InitRoute(e)
	e.Start(":" + os.Getenv("PORT"))
	fmt.Println("Hello World!")
}

func loadEnv() {
	err := godotenv.Load()
	if err!= nil {
        panic("Error loading.env file")
    }


}
