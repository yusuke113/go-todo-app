package main

import (
	"fmt"
	"go_todo_app/app/models"
	// "go_todo_app/config"
	// "log"
)

func main() {
	fmt.Println(models.Db)

	// u := &models.User{}
	// u.Name = "test2"
	// u.Email = "test2@example.com"
	// u.Password = "testtest"
	// fmt.Println(u)

	// u.CreatedUser()

	// u, _ := models.GetUser(1)

	// fmt.Println(u)

	// u.Name = "Test2"
	// u.Email = "Test2@example.com"
	// u.UpdateUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// u.DeleteUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// user, _ := models.GetUser(2)
	// user.CreatedTodo("First Todo")
	
	// t, _ := models.GetTodo(1)
	// fmt.Println(t)
	
	// user, _ := models.GetUser(3)
	// user.CreatedTodo("Third Todo")Z

	// todos, _ := models.GetTodos()
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }

	// user2, _ := models.GetUser(2)
	// todos, _ := user2.GetTodosByUser()
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }

	t, _ := models.GetTodo(3)
	// t.Content = "Update Todo"
	// t.UpdateTodo()
	t.DeleteTodo()
}
