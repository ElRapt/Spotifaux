package users

import (
	"encoding/json"
	"middleware/example/internal/models"
	users "middleware/example/internal/services/users"
	"net/http"

	"github.com/sirupsen/logrus"
)

// GetUsers
// @Tags         users
// @Summary      Get users.
// @Description  Get users.
// @Success      200            {array}  models.User
// @Failure      500             "Something went wrong"
// @Router       /users [get]
func GetUsers(w http.ResponseWriter, _ *http.Request) {
	// calling service
	users, err := users.GetAllUsers()
	if err != nil {
		logrus.Errorf("error : %s", err.Error())
		customError, isCustom := err.(*models.CustomError)
		if isCustom {
			w.WriteHeader(customError.Code)
			body, _ := json.Marshal(customError)
			_, _ = w.Write(body)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(users)
	_, _ = w.Write(body)
	return
}