package artists

import (
	"encoding/json"
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
func GetArtists(w http.ResponseWriter, _ *http.Request) {
	artists, err := artists.GetAllArtists()
	if err != nil {
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

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(artists)
	_, _ = w.Write(body)
	return
}
