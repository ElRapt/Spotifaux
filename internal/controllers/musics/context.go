package musics

import (
	"context"
	"fmt"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

func Ctx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		musicId, err := uuid.FromString(chi.URLParam(r, "id"))
		if err != nil {
			logrus.Errorf("parsing error : %s", err.Error())
			customError := &models.CustomError{
				Message: fmt.Sprintf("cannot parse id (%s) as UUID", chi.URLParam(r, "id")),
				Code:    http.StatusUnprocessableEntity,
			}
			w.WriteHeader(http.StatusInternalServerError)
			helpers.RespondWithFormat(w, r, customError)
		}

		ctx := context.WithValue(r.Context(), "musicId", musicId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
