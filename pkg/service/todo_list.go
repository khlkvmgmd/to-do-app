package service

import (
	todo "github.com/khlkvmgmd/to-do-app"
	"github.com/khlkvmgmd/to-do-app/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func newTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId int, list todo.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}