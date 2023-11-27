package genres

import (
	"encoding/json"
	"middleware/example/internal/models"
	"middleware/example/internal/services/genres"
	"net/http"

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
	ctx := r.Context()
	genreId, _ := ctx.Value("genreId").(uuid.UUID)

	var newGenre models.Genre
	err := json.NewDecoder(r.Body).Decode(&newGenre)
	if err != nil {
		logrus.Errorf("error : %s", err.Error())
		customError := &models.CustomError{
			Message: "cannot parse body",
			Code:    http.StatusUnprocessableEntity,
		}
		w.WriteHeader(customError.Code)
		body, _ := json.Marshal(customError)
		_, _ = w.Write(body)
		return
	}

	err = genres.PutGenre(genreId, newGenre)
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
	body, _ := json.Marshal("Genre updated")
	_, _ = w.Write(body)
	return
}
