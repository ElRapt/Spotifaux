package ratings

import (
	"encoding/json"
	"middleware/example/internal/models"
	"middleware/example/internal/services/ratings"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

// PutMusicRating
// @Tags         ratings
// @Summary      Modify a music rating.
// @Description  Modify a music rating.
// @Param        music_id        path      	string  				true  	"Music UUID formatted ID"
// @Param        rating_id      path      	string  				true  	"Rating UUID formatted ID"
// @Param        ratingRequest  body  		models.RatingRequest 	true	"rating request"
// @Success      200            {object}  	models.Rating
// @Failure      404            "rating not found"
// @Failure      422            "Cannot parse id"
// @Failure      422            "rating must be between 0 and 5"
// @Failure      500            "Something went wrong"
// @Router       /musics/{music_id}/ratings/{rating_id} [put]
func PutMusicRating(w http.ResponseWriter, r *http.Request) {
	var ratingRequest models.RatingRequest

	ctx := r.Context()
	musicID, _ := ctx.Value("musicID").(uuid.UUID)
	ratingID, _ := ctx.Value("ratingID").(uuid.UUID)

	err := json.NewDecoder(r.Body).Decode(&ratingRequest)
	if err != nil {
		logrus.Errorf("Data decoding error : %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	rating, err := ratings.ModifyMusicRating(musicID, ratingID, ratingRequest)
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
