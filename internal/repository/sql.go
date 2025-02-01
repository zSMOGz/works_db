package repository

import (
	"database/sql"
)

const (
	queryInsertUser = "INSERT INTO users (id, name) VALUES ($1, $2)"
	querySelectUser = "SELECT id, name FROM users WHERE id = $1"
	queryUpdateUser = "UPDATE users SET name = $1 WHERE id = $2"
	queryDeleteUser = "DELETE FROM users WHERE id = $1"
)

type SQLRepository struct {
	db *sql.DB
}

func NewSQLRepository(db *sql.DB) *SQLRepository {
	return &SQLRepository{db: db}
}

func (r *SQLRepository) Create(id int, name string) error {
	_, err := r.db.Exec(queryInsertUser, id, name)
	return err
}

func (r *SQLRepository) Get(id int) (int, string, error) {
	var userID int
	var userName string
	err := r.db.QueryRow(querySelectUser, id).Scan(&userID, &userName)
	return userID, userName, err
}

func (r *SQLRepository) Update(id int, newName string) error {
	_, err := r.db.Exec(queryUpdateUser, newName, id)
	return err
}

func (r *SQLRepository) Delete(id int) error {
	_, err := r.db.Exec(queryDeleteUser, id)
	return err
}
