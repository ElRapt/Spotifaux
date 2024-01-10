package musics

import (
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
	"middleware/example/internal/services/musics"
	"net/http"

	"github.com/sirupsen/logrus"
)

// GetMusics
// @Tags         musics
// @Summary      Get musics.
// @Description  Get musics.
// @Success      200            {array}  models.Music
// @Failure      500             "Something went wrong"
// @Router       /musics [get]
func GetMusics(w http.ResponseWriter, r *http.Request) {
	musics, err := musics.GetAllMusics()
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

	helpers.RespondWithFormat(w, r, musics)
}
