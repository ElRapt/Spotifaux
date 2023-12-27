package albums

import (
	"encoding/json"
	"encoding/xml"
	"middleware/example/internal/helpers"
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

	switch r.Header.Get("Content-Type") {
	case "application/xml":
		err := xml.NewDecoder(r.Body).Decode(&newAlbum)
		if err != nil {
			logrus.Errorf("error : %s", err.Error())
			customError := &models.CustomError{
				Message: "cannot parse body as XML",
				Code:    http.StatusUnprocessableEntity,
			}
			helpers.RespondWithFormat(w, r, customError)
			return
		}
	default:
		err := json.NewDecoder(r.Body).Decode(&newAlbum)
		if err != nil {
			logrus.Errorf("error : %s", err.Error())
			customError := &models.CustomError{
				Message: "cannot parse body as JSON",
				Code:    http.StatusUnprocessableEntity,
			}
			helpers.RespondWithFormat(w, r, customError)
			return
		}
	}

	albumId, err := albums.PostAlbum(newAlbum)
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

	album, err := albums.GetAlbumById(albumId)
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

	helpers.RespondWithFormat(w, r, album)
}
