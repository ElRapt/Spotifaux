package users

import (
	"database/sql"
	"errors"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"

	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

func GetAllUsers() ([]models.User, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM users")
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	users := []models.User{}
	for rows.Next() {
		var data models.User
		err = rows.Scan(
			&data.Id,
			&data.Username,
			&data.Email,
			&data.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, data)
	}

	_ = rows.Close()

	return users, err
}

func GetUserById(userId uuid.UUID) (models.User, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return models.User{}, err
	}
	defer db.Close()

	var user models.User

	err = db.QueryRow("SELECT id, username, email, created_at FROM users WHERE id = $1", userId).Scan(
		&user.Id,
		&user.Username,
		&user.Email,
		&user.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			// This error is tricky to test, got to have a valid uuid but non-existent in the DB
			return models.User{}, &models.CustomError{
				Message: "User not found",
				Code:    http.StatusNotFound,
			}
		}
		return models.User{}, err
	}

	return user, nil
}

func SaveUser(user models.User) error {
	db, err := helpers.OpenDB()
	if err != nil {
		logrus.Errorf("error opening database connection: %s", err.Error())
		return &models.CustomError{
			Message: "Error opening database connection",
			Code:    http.StatusInternalServerError,
		}
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO users (id, username, email) VALUES ($1, $2, $3)",
		user.Id.String(), user.Username, user.Email)
	if err != nil {
		logrus.Errorf("error inserting user into the database: %s", err.Error())
		return &models.CustomError{
			Message: "Error inserting user into the database",
			Code:    http.StatusInternalServerError,
		}
	}

	return nil
}

func UpdateUser(userId uuid.UUID, username string, email string) (models.User, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return models.User{}, err
	}
	defer db.Close()

	var updateQuery string
	var args []interface{}

	if username != "" && email != "" {
		updateQuery = "UPDATE users SET username = $1, email = $2 WHERE id = $3"
		args = append(args, username, email, userId)
	} else if username != "" {
		updateQuery = "UPDATE users SET username = $1 WHERE id = $2"
		args = append(args, username, userId)
	} else if email != "" {
		updateQuery = "UPDATE users SET email = $1 WHERE id = $2"
		args = append(args, email, userId)
	} else {
		return models.User{}, errors.New("no fields to update")
	}

	_, err = db.Exec(updateQuery, args...)
	if err != nil {
		return models.User{}, err
	}

	var user models.User
	err = db.QueryRow("SELECT id, username, email, created_at FROM users WHERE id = $1", userId).Scan(
		&user.Id,
		&user.Username,
		&user.Email,
		&user.CreatedAt,
	)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func DeleteUser(userId uuid.UUID) (models.User, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return models.User{}, err
	}
	defer db.Close()

	var user models.User
	err = db.QueryRow("SELECT id, username, email, created_at FROM users WHERE id = $1", userId).Scan(
		&user.Id,
		&user.Username,
		&user.Email,
		&user.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, &models.CustomError{
				Message: "User not found",
				Code:    http.StatusNotFound,
			}
		}
		return models.User{}, err
	}

	_, err = db.Exec("DELETE FROM users WHERE id = $1", userId)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
