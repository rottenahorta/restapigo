package repo

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	todo "github.com/rottenahorta/restapi101"
)

type TodoItemPostgres struct {
	db *sqlx.DB
}

func NewTodoItemPostgres(db *sqlx.DB) *TodoItemPostgres {
	return &TodoItemPostgres{db: db}
}

func (r *TodoItemPostgres) Create(uid, lid int, i todo.TodoItem) (int, error) {
	t, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var itemid int
	q := fmt.Sprintf("INSERT INTO %s (title, description, done) VALUES ($1,$2,$3) RETURNING id", todoItemsTable)
	row := t.QueryRow(q, i.Title, i.Description, i.Done)
	if err := row.Scan(&itemid); err != nil { 
		t.Rollback()
		return 0, err
	}
	q = fmt.Sprintf("INSERT INTO %s (item_id, list_id) VALUES ($1,$2)", listsItemsTable)
	_, err = t.Exec(q, itemid, lid)
	if err != nil {
		t.Rollback()
		return 0, err
	}
	return itemid, t.Commit()
}