package main

import (
	"fmt"
	"go_todo_app/app/controllers"
	"go_todo_app/app/models"
	// "go_todo_app/config"
	// "log"
)

func main() {
	fmt.Println(models.Db)

	controllers.StartMainServer()
}
