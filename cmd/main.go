package main

import (
	goTodoApp "go-todo-app"
	"log"
)

func main() {
	server := new(goTodoApp.Server)
	if err := server.Start("8100"); err != nil {
		log.Fatalf("Server start error: %s", err)
	}
}
