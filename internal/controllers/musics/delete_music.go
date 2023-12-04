package musics

import (
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
	"middleware/example/internal/repositories/musics"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

// DeleteMusic
// @Tags         musics
// @Summary      Delete a music.
// @Description  Delete a music.
// @Param        id           	path      string  true  "Music UUID formatted ID"
// @Success      200            {object}  string
// @Failure      422            "Cannot parse id"
// @Failure      500            "Something went wrong"
// @Router       /musics/{id} [delete]
func DeleteMusic(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	musicId, _ := ctx.Value("musicId").(uuid.UUID)

	err := musics.DeleteMusic(musicId)
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

	helpers.RespondWithFormat(w, r, "Music deleted")
}
