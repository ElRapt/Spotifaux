package ratings

import (
	"encoding/json"
	"middleware/example/internal/models"
	"middleware/example/internal/services/ratings"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

// GetMusicRating
// @Tags         ratings
// @Summary      Get a music rating.
// @Description  Get a music rating.
// @Param        music_id        path      string  true  "Music UUID formatted ID"
// @Param        rating_id      path      string  true   "Rating UUID formatted ID"
// @Success      200            {object}  models.Rating
// @Failure      404            "Rating not found"
// @Failure      422            "Cannot parse id"
// @Failure      500            "Something went wrong"
// @Router       /musics/{music_id}/ratings/{rating_id} [get]
func GetMusicRating(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	musicID, _ := ctx.Value("musicID").(uuid.UUID)
	ratingID, _ := ctx.Value("ratingID").(uuid.UUID)

	rating, err := ratings.GetMusicRating(musicID, ratingID)
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

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(rating)
	_, _ = w.Write(body)
	return
}
