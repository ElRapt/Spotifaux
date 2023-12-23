package users

import (
	"middleware/example2/internal/models"
	repository "middleware/example2/internal/repositories/users"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

func GetAllUsers() ([]models.User, error) {
	var err error
	users, err := repository.GetAllUsers()
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

func CreateUser(partialUser models.User) (models.User, error) {

	userId, err := uuid.NewV4()
	if err != nil {
		logrus.Fatalf("failed to generate UUID: %s", err.Error())
	}

	newUser := models.User{
		Id:       userId,
		Username: partialUser.Username,
		Email:    partialUser.Email,
	}

	err2 := repository.CreateUser(newUser)
	if err2 != nil {
		logrus.Errorf("error creating user: %s", err.Error())
		return models.User{}, err2
	}

	return newUser, nil
}

func UpdateUser(partialUser models.User) (models.User, error) {

	var user models.User
	var err2 error

	user, err2 = repository.UpdateUser(partialUser.Id, partialUser.Username, partialUser.Email)

	if err2 != nil {
		logrus.Errorf("error updating user: %s", err2.Error())
		return models.User{}, &models.CustomError{
			Message: "Error updating user",
			Code:    500,
		}
	}

	return user, nil
}

func DeleteUser(userId uuid.UUID) (models.User, error) {
	var err error

	user, err := repository.DeleteUser(userId)

	if err != nil {
		logrus.Errorf("error deleting user : %s", err.Error())
		return models.User{}, &models.CustomError{
			Message: "Something went wrong (user)",
			Code:    500,
		}
	}

	return user, nil
}
