package main

import (
	"log"
	todo "github.com/khlkvmgmd/to-do-app"
	"github.com/khlkvmgmd/to-do-app/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRotes()); err != nil{
		log.Fatalf("error: %s", err.Error())
	}
}