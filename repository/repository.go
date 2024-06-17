package repository

import (
	"database/sql"
	"encoding/json"
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

func (r Repository) GetUserByEmail(value string) (*User, error) {
	stat, err := r.db.Prepare("SELECT * FROM users u WHERE LOWER(email) = LOWER(?)")

	if err != nil {
		return nil, err
	}

	rows, err := stat.Query(value)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	exists := rows.Next()

	if !exists {
		return nil, nil
	}

	user := User{}

	err = rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r Repository) GetMuscleGroupsWithExercises() ([]dtos.MuscleGroupLookupResponse, error) {
	rows, err := r.db.Query(`
    SELECT 
        mg.id,
        mg.name,
        JSON_GROUP_ARRAY(
            JSON_OBJECT(
                'id', e.id,
                'label', e.name
            )
        ) AS exercises
    FROM 
        muscle_groups mg
    LEFT JOIN 
        exercises e ON mg.id = e.muscle_group_id
    GROUP BY 
        mg.id;
`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var results []dtos.MuscleGroupLookupResponse

	for rows.Next() {
		result := dtos.MuscleGroupLookupResponse{}
		var exercisesJSON []byte

		err = rows.Scan(&result.Id, &result.Label, &exercisesJSON)

		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(exercisesJSON, &result.Exercises)

		if err != nil {
			return nil, err
		}

		results = append(results, result)
	}

	return results, nil
}
