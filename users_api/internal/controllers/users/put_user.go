package users

import (
	"encoding/json"
	"io"
	"middleware/example2/internal/helpers"
	"middleware/example2/internal/models"
	users "middleware/example2/internal/services/users"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

// UpdateUser
// @Tags         users
// @Summary      Update user.
// @Description  Update an existing user's information.
// @Accept       json
// @Produce      json
// @Produce      xml
// @Param        id    path      string                true  "User ID"
// @Param        user  body      models.User  true  "User Data"
// @Success      200            "User Updated"
// @Failure      500            "Something went wrong"
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

	var requestBody map[string]interface{}
	if err := json.Unmarshal([]byte(bodyStr), &requestBody); err != nil {
		logrus.Errorf("error decoding JSON: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	username, okUsername := requestBody["username"].(string)
	if !okUsername {
		username = ""
	}
	email, okEmail := requestBody["email"].(string)
	if !okEmail {
		email = ""
	}

	if !okEmail && !okUsername {
		w.WriteHeader(http.StatusBadRequest)
		helpers.RespondWithFormat(w, r, map[string]string{"error": "username and email fields are both missing or not a string"})
		return
	}

	newUser := models.User{
		Id:       userId,
		Username: username,
		Email:    email,
	}

	updatedUser, err := users.UpdateUser(newUser)
	if err != nil {
		logrus.Errorf("error updating user: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	helpers.RespondWithFormat(w, r, updatedUser)
}
