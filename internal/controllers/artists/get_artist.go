package artists

import (
	"encoding/json"
	"middleware/example/internal/models"
	"middleware/example/internal/repositories/artists"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

// GetArtist
// @Tags         artists
// @Summary      Get an artist.
// @Description  Get an artist.
// @Param        id           	path      string  true  "Artist UUID formatted ID"
// @Success      200            {object}  models.Artist
// @Failure      422            "Cannot parse id"
// @Failure      500            "Something went wrong"
// @Router       /artists/{id} [get]
func GetArtist(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	artistId, _ := ctx.Value("artistId").(uuid.UUID)

	artist, err := artists.GetArtistById(artistId)
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
	body, _ := json.Marshal(artist)
	_, _ = w.Write(body)
	return
}
