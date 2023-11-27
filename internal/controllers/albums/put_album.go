package albums

import (
	"encoding/json"
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

	err = albums.PutAlbum(albumId, newAlbum)
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
	body, _ := json.Marshal("Album updated")
	_, _ = w.Write(body)
	return
}
