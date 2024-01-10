package artists

import (
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
	"middleware/example/internal/services/artists"
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
			helpers.RespondWithFormat(w, r, customError)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			helpers.RespondWithFormat(w, r, "Something went wrong")
		}
		return
	}

	helpers.RespondWithFormat(w, r, artist)
}
