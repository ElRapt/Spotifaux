package albums

import (
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
	"middleware/example/internal/services/albums"
	"net/http"

	"github.com/sirupsen/logrus"
)

// GetAlbums
// @Tags         albums
// @Summary      Get albums.
// @Description  Get albums.
// @Success      200            {array}  models.Album
// @Failure      500             "Something went wrong"
// @Router       /albums [get]
func GetAlbums(w http.ResponseWriter, r *http.Request) {
	albums, err := albums.GetAllAlbums()
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

	helpers.RespondWithFormat(w, r, albums)
}
