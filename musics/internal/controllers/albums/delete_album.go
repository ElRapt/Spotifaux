package albums

import (
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
	"middleware/example/internal/repositories/albums"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

// DeleteAlbum
// @Tags         albums
// @Summary      Delete an Album.
// @Description  Delete an Album.
// @Param        id           	path      string  true  "Album UUID formatted ID"
// @Success      200            {object}  string
// @Failure      422            "Cannot parse id"
// @Failure      500            "Something went wrong"
// @Router       /albums/{id} [delete]
func DeleteAlbum(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	albumId, _ := ctx.Value("albumId").(uuid.UUID)

	err := albums.DeleteAlbum(albumId)
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

	helpers.RespondWithFormat(w, r, "Album deleted")
}
