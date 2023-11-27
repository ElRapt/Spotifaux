package genres

import (
	"encoding/json"
	"middleware/example/internal/models"
	"middleware/example/internal/repositories/genres"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

// DeleteGenre
// @Tags         genres
// @Summary      Delete a Genre.
// @Description  Delete a Genre.
// @Param        id           	path      string  true  "Genre UUID formatted ID"
// @Success      200            {object}  string
// @Failure      422            "Cannot parse id"
// @Failure      500            "Something went wrong"
// @Router       /genres/{id} [delete]
func DeleteGenre(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	genreId, _ := ctx.Value("genreId").(uuid.UUID)

	err := genres.DeleteGenre(genreId)
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
	body, _ := json.Marshal("Genre deleted")
	_, _ = w.Write(body)
	return
}
