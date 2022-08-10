package service

import ("github.com/rottenahorta/restapi101/pkg/repo"
"github.com/rottenahorta/restapi101" )

type Auth interface {
	CreateUser(u todo.User) (int,error)
	GenerateToken(un, p string) (string, error)
	ParseJWT(t string) (int, error)
}

type TodoList interface {
	CreateList(uid int, l todo.TodoList) (int,error)
	GetAllLists(uid int) ([]todo.TodoList, error)
	GetList(uid, lid int) (todo.TodoList, error)
	DeleteList(uid, lid int) (error)
	UpdateList(uid, lid int, l todo.UpdateListInput) (error)
}

type TodoItem interface {
	Create(uid, lid int, i todo.TodoItem) (int, error)
}

type Service struct {
	Auth
	TodoList
	TodoItem
}

func NewService(r *repo.Repo) *Service {
	return &Service{
		Auth: NewAuthService(r.Auth),
		TodoList: NewTodoListService(r.TodoList),
		TodoItem: NewTodoItemService(r.TodoItem),
	}
}