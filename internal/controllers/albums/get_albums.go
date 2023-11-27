package albums

import (
	"encoding/json"
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
func GetAlbums(w http.ResponseWriter, _ *http.Request) {
	albums, err := albums.GetAllAlbums()
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
	body, _ := json.Marshal(albums)
	_, _ = w.Write(body)
	return
}
