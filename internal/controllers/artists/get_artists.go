package artists

import (
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
	"middleware/example/internal/services/artists"
	"net/http"

	"github.com/sirupsen/logrus"
)

// GetArtists
// @Tags         artists
// @Summary      Get artists.
// @Description  Get artists.
// @Success      200            {array}  models.Artist
// @Failure      500             "Something went wrong"
// @Router       /artists [get]
func GetArtists(w http.ResponseWriter, r *http.Request) {
	artists, err := artists.GetAllArtists()
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

	helpers.RespondWithFormat(w, r, artists)
}
