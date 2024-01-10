package ratings

import (
	"encoding/json"
	"middleware/example/internal/models"
	"middleware/example/internal/services/ratings"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

// DeleteMusicRating
// @Tags         ratings
// @Summary      Delete a music rating.
// @Description  Delete a music rating.
// @Param        music_id        path      string  true  "Music UUID formatted ID"
// @Param        rating_id       path      string  true  "Rating UUID formatted ID"
// @Success      204            "No Content"
// @Failure      422            "Cannot parse id"
// @Failure      500            "Something went wrong"
// @Router       /musics/{music_id}/ratings/{rating_id} [delete]
func DeleteMusicRating(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	musicID, _ := ctx.Value("musicID").(uuid.UUID)
	ratingID, _ := ctx.Value("ratingID").(uuid.UUID)

	err := ratings.DeleteMusicRating(musicID, ratingID)
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

	w.WriteHeader(http.StatusNoContent)
	return
}
