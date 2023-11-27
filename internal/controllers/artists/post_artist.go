package artists

import (
	"encoding/json"
	"middleware/example/internal/models"
	"middleware/example/internal/services/artists"
	"net/http"

	"github.com/sirupsen/logrus"
)

// PostArtist
// @Tags         artists
// @Summary      Create an Artist.
// @Description  Create an Artist.
// @Param        body         	body      string  true  "Artist object"
// @Success      200            {array}  models.Artist
// @Failure      500             "Something went wrong"
// @Router       /artists [post]
func PostArtist(w http.ResponseWriter, r *http.Request) {
	var newArtist models.Artist
	err := json.NewDecoder(r.Body).Decode(&newArtist)
	if err != nil {
		logrus.Errorf("error : %s", err.Error())
		customError := &models.CustomError{
			Message: "cannot parse body",
			Code:    http.StatusUnprocessableEntity,
		}
		w.WriteHeader(customError.Code)
		body, _ := json.Marshal(customError)
		_, _ = w.Write(body)
		return
	}

	ArtistId, err := artists.PostArtist(newArtist)
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

	Artist, err := artists.GetArtistById(ArtistId)
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
	body, _ := json.Marshal(Artist)
	_, _ = w.Write(body)
	return
}
