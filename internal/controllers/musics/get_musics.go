package musics

import (
	"encoding/json"
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
func GetMusics(w http.ResponseWriter, _ *http.Request) {
	musics, err := musics.GetAllMusics()
	if err != nil {
		logrus.Errorf("error : %s", err.Error())
		customError, isCustom := err.(*models.CustomError)
		if isCustom {
			// writing http code in header
			w.WriteHeader(customError.Code)
			// writing error message in body
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
