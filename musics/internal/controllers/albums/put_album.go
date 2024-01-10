package albums

import (
	"encoding/json"
	"encoding/xml"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
	"middleware/example/internal/services/albums"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

// PutAlbum
// @Tags         albums
// @Summary      Update an Album.
// @Description  Update an Album.
// @Param        id           	path      string  true  "Album UUID formatted ID"
// @Param        body         	body      string  true  "Album object"
// @Success      200            {object}  string
// @Failure      422            "Cannot parse id"
// @Failure      500            "Something went wrong"
// @Router       /albums/{id} [put]
func PutAlbum(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	albumId, _ := ctx.Value("albumId").(uuid.UUID)

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

	err := albums.PutAlbum(albumId, newAlbum)
	if err != nil {
		logrus.Errorf("error : %s", err.Error())
		customError, isCustom := err.(*models.CustomError)
		if isCustom {
			helpers.RespondWithFormat(w, r, customError)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			helpers.RespondWithFormat(w, r, "Something went wrong")
		}
		return
	}

	helpers.RespondWithFormat(w, r, "Album updated successfully")
}
