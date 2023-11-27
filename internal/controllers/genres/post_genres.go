package genres

import (
	"encoding/json"
	"middleware/example/internal/models"
	"middleware/example/internal/services/genres"
	"net/http"

	"github.com/sirupsen/logrus"
)

// PostGenre
// @Tags         genres
// @Summary      Create a Genre.
// @Description  Create a Genre.
// @Param        body         	body      string  true  "Genre object"
// @Success      200            {array}  models.Genre
// @Failure      500             "Something went wrong"
// @Router       /genres [post]
func PostGenre(w http.ResponseWriter, r *http.Request) {
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

	genreId, err := genres.PostGenre(newGenre)
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

	Genre, err := genres.GetGenreById(genreId)
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
	body, _ := json.Marshal(Genre)
	_, _ = w.Write(body)
	return
}
