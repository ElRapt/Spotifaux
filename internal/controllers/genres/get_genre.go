package genres

import (
	"encoding/json"
	"middleware/example/internal/models"
	"middleware/example/internal/services/genres"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

// GetGenre
// @Tags         genres
// @Summary      Get a genre.
// @Description  Get a genre.
// @Param        id           	path      string  true  "genre UUID formatted ID"
// @Success      200            {object}  models.genre
// @Failure      422            "Cannot parse id"
// @Failure      500            "Something went wrong"
// @Router       /genres/{id} [get]
func GetGenre(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	genreId, _ := ctx.Value("genreId").(uuid.UUID)

	genre, err := genres.GetGenreById(genreId)
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
	body, _ := json.Marshal(genre)
	_, _ = w.Write(body)
	return
}
