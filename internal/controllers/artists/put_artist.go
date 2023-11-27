package artists

import (
	"encoding/json"
	"middleware/example/internal/models"
	"middleware/example/internal/services/artists"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

// PutArtist
// @Tags         artists
// @Summary      Update an Artist.
// @Description  Update an Artist.
// @Param        id           	path      string  true  "Artist UUID formatted ID"
// @Param        body         	body      string  true  "Artist object"
// @Success      200            {object}  string
// @Failure      422            "Cannot parse id"
// @Failure      500            "Something went wrong"
// @Router       /artists/{id} [put]
func PutArtist(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	artistId, _ := ctx.Value("artistId").(uuid.UUID)

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

	err = artists.PutArtist(artistId, newArtist)
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
	body, _ := json.Marshal("Artist updated")
	_, _ = w.Write(body)
	return
}
