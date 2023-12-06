package collections

import (
	"encoding/json"
	"io"
	collections "middleware/example/internal/services/collections"
	"net/http"

	"github.com/sirupsen/logrus"
)

// CreateUser
// @Tags         users
// @Summary      Post user.
// @Description  Create a user.
// @Success      200            {array}  models.User
// @Failure      500             "Something went wrong"
// @Router       /users [post]
func CreateUser(w http.ResponseWriter, r *http.Request) {
	// Read the request body as a string
	body, err := io.ReadAll(r.Body)
	if err != nil {
		logrus.Errorf("error reading request body: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	bodyStr := string(body)

	newUser, err := collections.CreateUser(bodyStr)
	if err != nil {
		logrus.Errorf("error creating user: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	responseBody, _ := json.Marshal(newUser)
	_, _ = w.Write(responseBody)
}
