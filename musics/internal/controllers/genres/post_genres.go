package genres

import (
	"encoding/json"
	"encoding/xml"
	"middleware/example/internal/helpers"
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

	switch r.Header.Get("Content-Type") {
	case "application/xml":
		err := xml.NewDecoder(r.Body).Decode(&newGenre)
		if err != nil {
			logrus.Errorf("error : %s", err.Error())
			customError := &models.CustomError{
				Message: "cannot parse body as XML",
				Code:    http.StatusUnprocessableEntity,
			}
			helpers.RespondWithFormat(w, r, customError)
			return
		}
	default:
		err := json.NewDecoder(r.Body).Decode(&newGenre)
		if err != nil {
			logrus.Errorf("error : %s", err.Error())
			customError := &models.CustomError{
				Message: "cannot parse body as JSON",
				Code:    http.StatusUnprocessableEntity,
			}
			helpers.RespondWithFormat(w, r, customError)
			return
		}
	}

	genreId, err := genres.PostGenre(newGenre)
	if err != nil {
		logrus.Errorf("error : %s", err.Error())
		customError, isCustom := err.(*models.CustomError)
		if isCustom {
			w.WriteHeader(customError.Code)
			helpers.RespondWithFormat(w, r, customError)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			helpers.RespondWithFormat(w, r, "Something went wrong")
		}
		return
	}

	genre, err := genres.GetGenreById(genreId)
	if err != nil {
		logrus.Errorf("error : %s", err.Error())
		customError, isCustom := err.(*models.CustomError)
		if isCustom {
			w.WriteHeader(customError.Code)
			helpers.RespondWithFormat(w, r, customError)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			helpers.RespondWithFormat(w, r, "Something went wrong")
		}
		return
	}

	helpers.RespondWithFormat(w, r, genre)
}
