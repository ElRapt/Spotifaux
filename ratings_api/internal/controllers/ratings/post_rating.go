package ratings

import (
	"encoding/json"
	"encoding/xml"
	"middleware/example3/internal/helpers"
	"middleware/example3/internal/models"
	"middleware/example3/internal/services/ratings"
	"net/http"

	"github.com/sirupsen/logrus"
)

// PostRating
// @Tags         ratings
// @Summary      Create a Rating.
// @Description  Create a Rating.
// @Param        body         	body      string  true  "Rating object"
// @Success      200            {object}  models.Rating
// @Failure      500            "Something went wrong"
// @Router       /ratings [post]
func PostRating(w http.ResponseWriter, r *http.Request) {
	var newRating models.Rating

	switch r.Header.Get("Content-Type") {
	case "application/xml":
		err := xml.NewDecoder(r.Body).Decode(&newRating)
		if err != nil {
			logrus.Errorf("error: %s", err.Error())
			customError := &models.CustomError{
				Message: "cannot parse body as XML",
				Code:    http.StatusUnprocessableEntity,
			}
			helpers.RespondWithFormat(w, r, customError)
			return
		}
	default:
		err := json.NewDecoder(r.Body).Decode(&newRating)
		if err != nil {
			logrus.Errorf("error: %s", err.Error())
			customError := &models.CustomError{
				Message: "cannot parse body as JSON",
				Code:    http.StatusUnprocessableEntity,
			}
			helpers.RespondWithFormat(w, r, customError)
			return
		}
	}

	ratingId, err := ratings.PostRating(newRating)
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
