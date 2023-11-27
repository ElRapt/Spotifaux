package albums

import (
	"encoding/json"
	"middleware/example/internal/models"
	"middleware/example/internal/services/albums"
	"net/http"

	"github.com/sirupsen/logrus"
)

// PostAlbum
// @Tags         albums
// @Summary      Create an Album.
// @Description  Create an Album.
// @Param        body         	body      string  true  "Album object"
// @Success      200            {array}  models.Album
// @Failure      500             "Something went wrong"
// @Router       /albums [post]
func PostAlbum(w http.ResponseWriter, r *http.Request) {
	var newAlbum models.Album
	err := json.NewDecoder(r.Body).Decode(&newAlbum)
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

	albumId, err := albums.PostAlbum(newAlbum)
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

	Album, err := albums.GetAlbumById(albumId)
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
	body, _ := json.Marshal(Album)
	_, _ = w.Write(body)
	return
}
