package main

import (
	"middleware/example/internal/controllers/ratings"
	"middleware/example/internal/helpers"
	_ "middleware/example/internal/models"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

func main() {
	r := chi.NewRouter()

	// Route pour les notations (ratings)
	r.Route("/ratings", func(r chi.Router) {
		r.Get("/", ratings.GetRatings)
		r.Post("/", ratings.PostRating)
		r.Route("/{id}", func(r chi.Router) {
			r.Use(ratings.Ctx)
			r.Get("/", ratings.GetRating)
			r.Delete("/", ratings.DeleteRating)
			r.Put("/", ratings.PutRating)
		})
	})

	logrus.Info("[INFO] Web server started. Now listening on *:8080")
	logrus.Fatalln(http.ListenAndServe(":8080", r))
}

func init() {
	db, err := helpers.OpenDB()
	if err != nil {
		logrus.Fatalf("error while opening database: %s", err.Error())
	}
	schemes := []string{
		`
        CREATE TABLE IF NOT EXISTS Ratings (
            id CHAR(36) PRIMARY KEY,
            user_id CHAR(36) NOT NULL,
            song_id CHAR(36) NOT NULL,
            rating INT NOT NULL,
            comment TEXT
        );
        `,
	}
	for _, scheme := range schemes {
		if _, err := db.Exec(scheme); err != nil {
			logrus.Fatalln("Could not generate table! Error was: " + err.Error())
		}
	}
	helpers.CloseDB(db)
}
