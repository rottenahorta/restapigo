package repo

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	todo "github.com/rottenahorta/restapi101"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (r *TodoListPostgres) CreateList(uid int, l todo.TodoList) (int, error) {
	t, err := r.db.Begin() // init transaction
	if err != nil {
		return 0, err
	}

	var listid int
	q := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1,$2) RETURNING id", todoListsTable)
	row := t.QueryRow(q, l.Title, l.Description)
	if err := row.Scan(&listid); err != nil { // assignin id of list to listid var
		t.Rollback()
		return 0, err
	}
	q = fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1,$2)", userListsTable)
	_, err = t.Exec(q, uid, listid)
	if err != nil {
		t.Rollback()
		return 0, err
	}
	return listid, t.Commit()
}

func (r *TodoListPostgres) GetAllLists(uid int) ([]todo.TodoList, error) {
	var lists []todo.TodoList
	q := fmt.Sprintf("SELECT tl.* FROM %s tl INNER JOIN %s ul ON tl.id = ul.list_id WHERE ul.user_id = $1", todoListsTable, userListsTable)
	err := r.db.Select(&lists, q, uid)
	return lists, err
}

func (r *TodoListPostgres) GetList(uid, lid int) (todo.TodoList, error) {
	var list todo.TodoList
	q := fmt.Sprintf("SELECT tl.* FROM %s tl INNER JOIN %s ul ON tl.id = ul.list_id WHERE ul.list_id = $1 AND ul.user_id = $2", todoListsTable, userListsTable)
	err := r.db.Get(&list, q, lid, uid)
	return list, err
}

func (r *TodoListPostgres) DeleteList(uid, lid int) error {
	q := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id = ul.list_id AND ul.list_id = $1 AND ul.user_id = $2", todoListsTable, userListsTable)
	_, err := r.db.Exec(q, lid, uid)
	return err
}

func (r *TodoListPostgres) UpdateList(uid, lid int, l todo.UpdateListInput) error {
	values := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	if l.Title != nil {
		values = append(values, "title=$1")
		args = append(args, *l.Title)
		argId++
	} 
	if l.Description != nil {
		values = append(values, fmt.Sprintf("description=$%d",argId))
		args = append(args, *l.Description)
		argId++
	}
	qargs := strings.Join(values, ",")
	q := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.list_id AND ul.list_id = $%d AND ul.user_id = $%d",
		todoListsTable, qargs, userListsTable, argId, argId+1)
	args = append(args, lid, uid)
	_, err := r.db.Exec(q, args...)	
	return err
}
