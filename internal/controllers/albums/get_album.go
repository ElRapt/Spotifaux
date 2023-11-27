package albums

import (
	"encoding/json"
	"middleware/example/internal/models"
	"middleware/example/internal/services/albums"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

// Getalbum
// @Tags         albums
// @Summary      Get an album.
// @Description  Get an album.
// @Param        id           	path      string  true  "Album UUID formatted ID"
// @Success      200            {object}  models.Album
// @Failure      422            "Cannot parse id"
// @Failure      500            "Something went wrong"
// @Router       /albums/{id} [get]
func GetAlbum(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	albumId, _ := ctx.Value("albumId").(uuid.UUID)

	album, err := albums.GetAlbumById(albumId)
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
	body, _ := json.Marshal(album)
	_, _ = w.Write(body)
	return
}
