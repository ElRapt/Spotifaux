package musics

import (
	"encoding/json"
	"encoding/xml"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
	"middleware/example/internal/services/musics"
	"net/http"

	"github.com/sirupsen/logrus"
)

// PostMusic
// @Tags         musics
// @Summary      Create a music.
// @Description  Create a music.
// @Param        body         	body      string  true  "Music object"
// @Success      201            {array}  models.Music
// @Failure      500             "Something went wrong"
// @Router       /musics [post]
func PostMusic(w http.ResponseWriter, r *http.Request) {
	var newMusic models.Music

	switch r.Header.Get("Content-Type") {
	case "application/xml":
		err := xml.NewDecoder(r.Body).Decode(&newMusic)
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
		err := json.NewDecoder(r.Body).Decode(&newMusic)
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

	musicId, err := musics.PostMusic(newMusic)
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

	music, err := musics.GetMusicById(musicId)
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

	helpers.RespondWithFormat(w, r, music)
}
