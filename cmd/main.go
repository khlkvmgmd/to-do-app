package main

import (
	"log"

	todo "github.com/khlkvmgmd/to-do-app"
)

func main() {
	srv := new(todo.Server)
	if err := srv.Run("8000"); err != nil{
		log.Fatalf("error: %s", err.Error())
	}
}