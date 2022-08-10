package repo

import ("github.com/jmoiron/sqlx"
"github.com/rottenahorta/restapi101")

type Auth interface {
	CreateUser (u todo.User) (int, error) 
	GetUserCred (username, password string) (todo.User, error)
}

type TodoList interface {
	CreateList(uid int, list todo.TodoList) (int,error)
	GetAllLists(uid int) ([]todo.TodoList, error)
	GetList(uid, lid int) (todo.TodoList, error)
	DeleteList(uid, lid int) (error)
	UpdateList(uid, lid int, l todo.UpdateListInput) (error)
}

type TodoItem interface {
	Create(uid, lid int, i todo.TodoItem) (int, error)
}

type Repo struct {
	Auth
	TodoList
	TodoItem
}

func NewRepo(db *sqlx.DB) *Repo {
	return &Repo{
		Auth: NewAuthPostgres(db),
		TodoList: NewTodoListPostgres(db),
		TodoItem: NewTodoItemPostgres(db),
	}
}