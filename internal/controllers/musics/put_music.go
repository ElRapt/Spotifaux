package musics

import (
	"encoding/json"
	"middleware/example/internal/models"
	"middleware/example/internal/services/musics"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

// PutMusic
// @Tags         musics
// @Summary      Update a music.
// @Description  Update a music.
// @Param        id           	path      string  true  "Music UUID formatted ID"
// @Param        body         	body      string  true  "Music object"
// @Success      200            {object}  string
// @Failure      422            "Cannot parse id"
// @Failure      500            "Something went wrong"
// @Router       /musics/{id} [put]
func PutMusic(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	musicId, _ := ctx.Value("musicId").(uuid.UUID)

	var newMusic models.Music
	err := json.NewDecoder(r.Body).Decode(&newMusic)
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

	err = musics.PutMusic(musicId, newMusic)
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
	body, _ := json.Marshal("music updated")
	_, _ = w.Write(body)
	return
}
