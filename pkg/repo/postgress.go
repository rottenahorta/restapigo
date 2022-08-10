package repo

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	usersTable = "users"
	todoListsTable = "todo_lists"
	listsItemsTable = "lists_items"
	userListsTable = "user_lists"
	todoItemsTable = "todo_items"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgressDB(c Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", c.Host, c.Port, c.Username, c.Password, c.DBName, c.SSLMode))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}