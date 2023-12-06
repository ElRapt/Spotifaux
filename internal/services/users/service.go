package collections

import (
	"middleware/example/internal/models"
	repository "middleware/example/internal/repositories/collections"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

func GetAllUsers() ([]models.User, error) {
	var err error
	// calling repository
	users, err := repository.GetAllUsers()
	// managing errors
	if err != nil {
		logrus.Errorf("error retrieving users : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong (users)",
			Code:    500,
		}
	}

	return users, nil
}

func GetUserById(userId uuid.UUID) (models.User, error) {
	var err error

	user, err := repository.GetUserById(userId)

	if err != nil {
		logrus.Errorf("error retrieving user : %s", err.Error())
		return models.User{}, &models.CustomError{
			Message: "Something went wrong (user)",
			Code:    500,
		}
	}

	return user, nil
}
