package main

import (
	"fmt"
	"go_todo_app/app/controllers"
	"go_todo_app/app/models"
)

func main() {
	fmt.Println(models.Db)

	controllers.StartMainServer()

}
