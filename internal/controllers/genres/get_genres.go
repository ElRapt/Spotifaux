package genres

import (
	"encoding/json"
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
func GetGenres(w http.ResponseWriter, _ *http.Request) {
	genres, err := genres.GetAllGenres()
	if err != nil {
		logrus.Errorf("error : %s", err.Error())
		customError, isCustom := err.(*models.CustomError)
		if isCustom {
			// writing http code in header
			w.WriteHeader(customError.Code)
			// writing error message in body
			body, _ := json.Marshal(customError)
			_, _ = w.Write(body)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(genres)
	_, _ = w.Write(body)
	return
}
