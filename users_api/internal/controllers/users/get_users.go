package users

import (
	"middleware/example2/internal/helpers"
	"middleware/example2/internal/models"
	users "middleware/example2/internal/services/users"
	"net/http"

	"github.com/sirupsen/logrus"
)

// GetUsers
// @Tags         users
// @Summary      Get users.
// @Description  Get users.
// @Produce      json
// @Produce      xml
// @Success      200            {array}  models.User
// @Failure      500             "Something went wrong"
// @Router       /users [get]
func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := users.GetAllUsers()
	if err != nil {
		logrus.Errorf("error : %s", err.Error())
		customError, isCustom := err.(*models.CustomError)
		if isCustom {
			w.WriteHeader(customError.Code)
			helpers.RespondWithFormat(w, r, customError)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			helpers.RespondWithFormat(w, r, map[string]string{"error": "Internal Server Error"})
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	helpers.RespondWithFormat(w, r, users)
	return
}
