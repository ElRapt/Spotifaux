package ratings

import (
	"encoding/json"
	"middleware/example/internal/models"
	"middleware/example/internal/services/ratings"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

// GetMusicRatings
// @Tags         ratings
// @Summary      Get music ratings.
// @Description  Get music ratings.
// @Param        music_id        path     string  true   "Music UUID formatted ID"
// @Success      200            {array}  models.Rating
// @Failure      422            "Cannot parse id"
// @Failure      500            "Something went wrong"
// @Router       /musics/{music_id}/ratings [get]
func GetMusicRatings(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	musicID, _ := ctx.Value("musicID").(uuid.UUID)

	// calling service
	ratings, err := ratings.GetAllRatingsForAMusic(musicID)
	if err != nil {
		// logging error
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

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(ratings)
	_, _ = w.Write(body)
	return
}
