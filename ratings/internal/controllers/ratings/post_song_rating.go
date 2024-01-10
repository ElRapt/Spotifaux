package ratings

import (
	"encoding/json"
	"middleware/example/internal/models"
	"middleware/example/internal/services/ratings"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

// PostMusicRating
// @Tags         ratings
// @Summary      Post a music rating.
// @Description  Post a music rating.
// @Param        music_id        path      	string  				true  	"Music UUID formatted ID"
// @Param        ratingRequest  body  		models.RatingRequest 	true	"rating request"
// @Success      200            {object}  	models.Rating
// @Failure      422            "Cannot parse id"
// @Failure      422            "missing fields"
// @Failure      422            "rating must be between 0 and 5"
// @Failure      500            "Something went wrong"
// @Router       /musics/{music_id}/ratings [post]
func PostMusicRating(w http.ResponseWriter, r *http.Request) {
	var ratingRequest models.RatingRequest

	ctx := r.Context()
	musicID, _ := ctx.Value("musicID").(uuid.UUID)

	err := json.NewDecoder(r.Body).Decode(&ratingRequest)
	if err != nil {
		logrus.Errorf("Data decoding error : %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	rating, err := ratings.AddMusicRating(musicID, ratingRequest)
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

	ratingURL := "/musics/" + rating.MusicID.String() + "/ratings/" + rating.Id.String()
	w.Header().Set("Location", ratingURL)

	w.WriteHeader(http.StatusCreated)
	body, _ := json.Marshal(rating)
	_, _ = w.Write(body)
	return
}
