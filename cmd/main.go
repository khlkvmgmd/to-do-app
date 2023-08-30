package main

import (
	"log"

	"github.com/khlkvmgmd/to-do-app"
	"github.com/khlkvmgmd/to-do-app/pkg/handler"
	"github.com/khlkvmgmd/to-do-app/pkg/repository"
	"github.com/khlkvmgmd/to-do-app/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)


	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRotes()); err != nil{
		log.Fatalf("error: %s", err.Error())
	}
}