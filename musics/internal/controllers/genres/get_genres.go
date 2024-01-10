package genres

import (
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
	"middleware/example/internal/services/genres"
	"net/http"

	"github.com/sirupsen/logrus"
)

// GetGenres
// @Tags         genres
// @Summary      Get genres.
// @Description  Get genres.
// @Success      200            {array}  models.Genre
// @Failure      500             "Something went wrong"
// @Router       /genres [get]
func GetGenres(w http.ResponseWriter, r *http.Request) {
	genres, err := genres.GetAllGenres()
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

	helpers.RespondWithFormat(w, r, genres)
}
