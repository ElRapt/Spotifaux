package users

import (
	"encoding/json"
	"io"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
	users "middleware/example/internal/services/users"
	"net/http"

	"github.com/sirupsen/logrus"
)

// CreateUser
// @Tags         users
// @Summary      Post user.
// @Description  Create a new user with provided information.
// @Accept       json
// @Produce      json
// @Produce      xml
// @Param        user  body      models.User  true  "User Data"
// @Success      201            "User Created"
// @Failure      400            "Bad Request"
// @Failure      409            "Conflict"
// @Failure      500            "Something went wrong"
// @Router       /users [post]
func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		logrus.Errorf("error reading request body: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	bodyStr := string(body)

	var requestBody map[string]interface{}

	if err := json.Unmarshal([]byte(bodyStr), &requestBody); err != nil {
		logrus.Errorf("error decoding JSON: %s", err.Error())
		return
	}

	username, ok := requestBody["username"].(string)
	if !ok {
		logrus.Error("username field is missing or not a string", err.Error())
		return
	}

	email, ok := requestBody["email"].(string)
	if !ok {
		logrus.Error("email field is missing or not a string", err.Error())
		return
	}

	newUser := models.User{
		Username: username,
		Email:    email,
	}

	newUserR, err := users.CreateUser(newUser)
	if err != nil {
		logrus.Errorf("error creating user: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	helpers.RespondWithFormat(w, r, newUserR)
}
