package users

import (
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
	users "middleware/example/internal/services/users"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

// GetUser
// @Tags         users
// @Summary      Get user.
// @Description  Get user.
// @Success      200            {array}  models.User
// @Failure      500             "Something went wrong"
// @Router       /users/{id} [get]
func GetUserById(w http.ResponseWriter, r *http.Request) {
	userIdStr := chi.URLParam(r, "id")
	userId, err := uuid.FromString(userIdStr)
	if err != nil {
		logrus.Errorf("invalid UUID: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := users.GetUserById(userId)
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
	helpers.RespondWithFormat(w, r, user)
	return
}
