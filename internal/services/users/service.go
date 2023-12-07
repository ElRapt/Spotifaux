package users

import (
	"encoding/json"
	"errors"
	"middleware/example/internal/models"
	repository "middleware/example/internal/repositories/users"

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

func CreateUser(body string) (models.User, error) {
	var requestBody map[string]interface{}

	if err := json.Unmarshal([]byte(body), &requestBody); err != nil {
		logrus.Errorf("error decoding JSON: %s", err.Error())
		return models.User{}, err
	}

	username, ok := requestBody["username"].(string)
	if !ok {
		return models.User{}, errors.New("username field is missing or not a string")
	}

	email, ok := requestBody["email"].(string)
	if !ok {
		return models.User{}, errors.New("email field is missing or not a string")
	}

	userId, err := uuid.NewV4()
	if err != nil {
		logrus.Fatalf("failed to generate UUID: %s", err.Error())
	}

	newUser := models.User{
		Id:       userId,
		Username: username,
		Email:    email,
	}

	err2 := repository.SaveUser(newUser)
	if err2 != nil {
		logrus.Errorf("error creating user: %s", err.Error())
		return models.User{}, err2
	}

	return newUser, nil
}

func UpdateUser(userId uuid.UUID, body string) (models.User, error) {

	var requestBody map[string]interface{}
	if err := json.Unmarshal([]byte(body), &requestBody); err != nil {
		logrus.Errorf("error decoding JSON: %s", err.Error())
		return models.User{}, err
	}

	username, okUsername := requestBody["username"].(string)
	email, okEmail := requestBody["email"].(string)

	if !okEmail && !okUsername {
		return models.User{}, errors.New("username and email fields are missing or not a string")
	}

	var user models.User
	var err2 error

	user, err2 = repository.UpdateUser(userId, username, email)

	if err2 != nil {
		logrus.Errorf("error updating user: %s", err2.Error())
		return models.User{}, &models.CustomError{
			Message: "Error updating user",
			Code:    500,
		}
	}

	return user, nil
}
