package users

import (
	"middleware/example2/internal/helpers"
	"middleware/example2/internal/models"
	users "middleware/example2/internal/services/users"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

// DeleteUser
// @Tags         users
// @Summary      Delete user.
// @Description  Delete a user by their unique ID.
// @Produce      json
// @Produce      xml
// @Param        id    path      string  true  "User ID"
// @Success      200            "User Deleted"
// @Failure      400            "Bad Request"
// @Failure      404            "User Not Found"
// @Failure      500            "Internal Server Error"
// @Router       /users/{id} [delete]
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userIdStr := chi.URLParam(r, "id")
	userId, err := uuid.FromString(userIdStr)
	if err != nil {
		logrus.Errorf("invalid UUID: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := users.DeleteUser(userId)
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
