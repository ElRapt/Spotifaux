package ratings

import (
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
	"middleware/example/internal/services/ratings"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

// GetRating
// @Tags         ratings
// @Summary      Get a rating.
// @Description  Get a rating.
// @Param        id           	path      string  true  "Rating UUID formatted ID"
// @Success      200            {object}  models.Rating
// @Failure      422            "Cannot parse id"
// @Failure      500            "Something went wrong"
// @Router       /ratings/{id} [get]
func GetRating(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ratingId, _ := ctx.Value("ratingId").(uuid.UUID)

	rating, err := ratings.GetRatingById(ratingId)
	if err != nil {
		logrus.Errorf("error: %s", err.Error())
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

	helpers.RespondWithFormat(w, r, rating)
}
