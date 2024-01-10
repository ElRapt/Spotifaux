package genres

import (
	"encoding/json"
	"io"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
	genres "middleware/example/internal/services/genres"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

// PutGenre
// @Tags         genres
// @Summary      Update a Genre.
// @Description  Update a Genre.
// @Param        id           	path      string  true  "Genre UUID formatted ID"
// @Param        body         	body      string  true  "Genre object"
// @Success      200            {object}  string
// @Failure      422            "Cannot parse id"
// @Failure      500            "Something went wrong"
// @Router       /genres/{id} [put]
func PutGenre(w http.ResponseWriter, r *http.Request) {
	genreIdStr := chi.URLParam(r, "id")
	genreId, err := uuid.FromString(genreIdStr)
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

	name, okName := requestBody["name"].(string)
	if !okName {
		w.WriteHeader(http.StatusBadRequest)
		helpers.RespondWithFormat(w, r, map[string]string{"error": "name field is missing or not a string"})
		return
	}

	newGenre := models.Genre{
		Id:   genreId,
		Name: name,
	}

	updatedGenre, err := genres.PutGenre(newGenre)
	if err != nil {
		logrus.Errorf("error updating Genre: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	helpers.RespondWithFormat(w, r, updatedGenre)
}
