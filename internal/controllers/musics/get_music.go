package musics

import (
	"encoding/json"
	"middleware/example/internal/models"
	"middleware/example/internal/repositories/musics"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

// GetMusic
// @Tags         musics
// @Summary      Get a music.
// @Description  Get a music.
// @Param        id           	path      string  true  "Music UUID formatted ID"
// @Success      200            {object}  models.Music
// @Failure      422            "Cannot parse id"
// @Failure      500            "Something went wrong"
// @Router       /musics/{id} [get]
func GetMusic(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	musicId, _ := ctx.Value("musicId").(uuid.UUID)

	musics, err := musics.GetMusicById(musicId)
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
	body, _ := json.Marshal(musics)
	_, _ = w.Write(body)
	return
}
