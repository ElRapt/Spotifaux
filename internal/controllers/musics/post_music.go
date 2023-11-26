package musics

import (
	"encoding/json"
	"middleware/example/internal/models"
	"middleware/example/internal/services/musics"
	"net/http"

	"github.com/sirupsen/logrus"
)

// PostMusic
// GetMusics
// @Tags         musics
// @Summary      Create a music.
// @Description  Create a music.
// @Param        body         	body      string  true  "Music object"
// @Success      200            {array}  models.Music
// @Failure      500             "Something went wrong"
// @Router       /musics [post]
func PostMusic(w http.ResponseWriter, r *http.Request) {
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

	musicId, err := musics.PostMusic(newMusic)
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

	music, err := musics.GetMusicById(musicId)
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
	body, _ := json.Marshal(music)
	_, _ = w.Write(body)
	return
}
