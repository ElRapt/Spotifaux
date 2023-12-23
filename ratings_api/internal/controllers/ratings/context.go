package ratings

import (
	"context"
	"fmt"
	"middleware/example3/internal/helpers"
	"middleware/example3/internal/models"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

func Ctx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ratingId, err := uuid.FromString(chi.URLParam(r, "id"))
		if err != nil {
			logrus.Errorf("parsing error: %s", err.Error())
			customError := &models.CustomError{
				Message: fmt.Sprintf("cannot parse id (%s) as UUID", chi.URLParam(r, "id")),
				Code:    http.StatusUnprocessableEntity,
			}
			w.WriteHeader(http.StatusInternalServerError)
			helpers.RespondWithFormat(w, r, customError)
			return // Ajouter un return pour arrêter l'exécution ici en cas d'erreur
		}

		ctx := context.WithValue(r.Context(), "ratingId", ratingId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
