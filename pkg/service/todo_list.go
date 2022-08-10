package service

import (
	todo "github.com/rottenahorta/restapi101"
	"github.com/rottenahorta/restapi101/pkg/repo"
)

type TodoListService struct {
	repo repo.TodoList
}

func NewTodoListService(r repo.TodoList) *TodoListService {
	return &TodoListService{repo:r}
}

func (s *TodoListService) CreateList(uid int, list todo.TodoList) (int,error) {
	return s.repo.CreateList(uid, list)
}

func (s *TodoListService) GetAllLists(uid int) ([]todo.TodoList,error) {
	return s.repo.GetAllLists(uid)
}

func (s *TodoListService) GetList(uid, lid int) (todo.TodoList,error) {
	return s.repo.GetList(uid, lid)
}

func (s *TodoListService) UpdateList(uid, lid int, l todo.UpdateListInput) (error) {
	if err := l.Validate(); err!=nil { return err }
	return s.repo.UpdateList(uid, lid, l)
}

func (s *TodoListService) DeleteList(uid, lid int) (error) {
	return s.repo.DeleteList(uid, lid)
}