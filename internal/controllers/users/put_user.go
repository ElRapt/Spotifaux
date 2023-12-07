package users

import (
	"io"
	"middleware/example/internal/helpers"
	users "middleware/example/internal/services/users"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

// UpdateUser
// @Tags         users
// @Summary      Modify user.
// @Description  Modify a user.
// @Success      201            {array}  models.User
// @Failure      204              "Something went wrong"
// @Router       /users/{id} [put]

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	userIdStr := chi.URLParam(r, "id")
	userId, err := uuid.FromString(userIdStr)
	if err != nil {
		logrus.Errorf("invalid UUID: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		logrus.Errorf("error reading request body: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	bodyStr := string(body)

	updatedUser, err := users.UpdateUser(userId, bodyStr)
	if err != nil {
		logrus.Errorf("error updating user: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	helpers.RespondWithFormat(w, r, updatedUser)
}
