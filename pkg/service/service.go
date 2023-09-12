package service

import (
	todo "github.com/khlkvmgmd/to-do-app"
	"github.com/khlkvmgmd/to-do-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoItem
	TodoList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthSevice(repos.Authorization),
	}
}