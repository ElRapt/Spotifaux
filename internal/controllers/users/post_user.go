package users

import (
	"io"
	"middleware/example/internal/helpers"
	users "middleware/example/internal/services/users"
	"net/http"

	"github.com/sirupsen/logrus"
)

// CreateUser
// @Tags         users
// @Summary      Post user.
// @Description  Create a user.
// @Success      201 User created           {array}  models.User
// @Failure      204 No content             "Something went wrong"
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

	newUser, err := users.CreateUser(bodyStr)
	if err != nil {
		logrus.Errorf("error creating user: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	helpers.RespondWithFormat(w, r, newUser)
}
