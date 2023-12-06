package collections

import (
	"database/sql"
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
