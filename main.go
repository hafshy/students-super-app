package main

import (
	"fmt"

	"github.com/hafshy/students-super-app/configs"
)

func main() {
	configs.InitDB()
	fmt.Println("Hello World!")
}
