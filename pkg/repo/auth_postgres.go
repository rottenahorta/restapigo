package repo

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	todo "github.com/rottenahorta/restapi101"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(u todo.User) (int, error) {
	var id int
	q := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", usersTable)
	row := r.db.QueryRow(q, u.Name, u.Username, u.Password) // generatin instruction for db
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUserCred(un, p string) (todo.User, error) {
	var u todo.User
	q := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&u, q, un, p)
	return u, err
}
