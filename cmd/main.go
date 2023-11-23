package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/controllers/collections"
	"middleware/example/internal/helpers"
	_ "middleware/example/internal/models"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	r.Route("/collections", func(r chi.Router) {
		r.Get("/", collections.GetCollections)
		r.Route("/{id}", func(r chi.Router) {
			r.Use(collections.Ctx)
			r.Get("/", collections.GetCollection)
		})
	})

	logrus.Info("[INFO] Web server started. Now listening on *:8080")
	logrus.Fatalln(http.ListenAndServe(":8080", r))
}

func init() {
	db, err := helpers.OpenDB()
	if err != nil {
		logrus.Fatalf("error while opening database : %s", err.Error())
	}
	schemes := []string{
		`
		CREATE TABLE IF NOT EXISTS Genre (
			id INTEGER PRIMARY KEY,
			name TEXT NOT NULL
		);
		
		CREATE TABLE IF NOT EXISTS Artist (
			id INTEGER PRIMARY KEY,
			name TEXT NOT NULL
		);
		
		CREATE TABLE IF NOT EXISTS Album (
			id INTEGER PRIMARY KEY,
			name TEXT NOT NULL,
			artistId INTEGER,
			FOREIGN KEY (artistId) REFERENCES Artist(id)
		);
		
		CREATE TABLE IF NOT EXISTS Music (
			id INTEGER PRIMARY KEY,
			title TEXT NOT NULL,
			genreId INTEGER,
			artistId INTEGER,
			albumId INTEGER,
			FOREIGN KEY (genreId) REFERENCES Genre(id),
			FOREIGN KEY (artistId) REFERENCES Artist(id),
			FOREIGN KEY (albumId) REFERENCES Album(id)
		);
		
		CREATE INDEX IF NOT EXISTS idx_artist_name ON Artist(name);
		CREATE INDEX IF NOT EXISTS idx_album_name ON Album(name);
		`,
	}
	for _, scheme := range schemes {
		if _, err := db.Exec(scheme); err != nil {
			logrus.Fatalln("Could not generate table ! Error was : " + err.Error())
		}
	}
	helpers.CloseDB(db)
}
