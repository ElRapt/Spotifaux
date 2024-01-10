package musics

import (
	"encoding/json"
	"encoding/xml"
	"middleware/example/internal/helpers"
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

	err := musics.PutMusic(musicId, newMusic)
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

	helpers.RespondWithFormat(w, r, "Music updated successfully")
}
