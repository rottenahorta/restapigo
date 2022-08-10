package service

import (
	todo "github.com/rottenahorta/restapi101"
	"github.com/rottenahorta/restapi101/pkg/repo"
)

type TodoItemService struct {
	repo repo.TodoItem
}

func NewTodoItemService(r repo.TodoItem) *TodoItemService {
	return &TodoItemService{repo:r}
}

func (s *TodoItemService) Create(uid, lid int, i todo.TodoItem) (int,error) {
	return s.repo.Create(uid, lid, i)
}