package repository

import (
	"database/sql"
	"fmt"

	"github.com/KhaledAlorayir/go-htmx-thinge/dtos"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return Repository{db: db}
}

func (r Repository) CreateUser(user dtos.CreateUserRequest) error {
	tx, err := r.db.Begin()

	if err != nil {
		return err
	}

	stat, err := tx.Prepare("INSERT INTO users (username, email, password) VALUES(?,?,?)")

	if err != nil {
		return err
	}

	defer stat.Close()

	_, err = stat.Exec(user.Username, user.Email, user.Password)

	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r Repository) checkIfRecordExists(table string, column string, value string) (bool, error) {
	stat, err := r.db.Prepare(fmt.Sprintf("SELECT 1 from %s WHERE LOWER(%s) = LOWER(?)", table, column))

	if err != nil {
		return false, err
	}

	rows, err := stat.Query(value)

	if err != nil {
		return false, err
	}

	defer rows.Close()

	hasResults := rows.Next()

	return hasResults, nil
}

func (r Repository) CheckIfUsernameExists(value string) (bool, error) {
	return r.checkIfRecordExists("users", "username", value)
}

func (r Repository) CheckIfEmailExists(value string) (bool, error) {
	return r.checkIfRecordExists("users", "email", value)
}
