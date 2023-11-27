package main

import (
	"middleware/example/internal/controllers/artists"
	"middleware/example/internal/controllers/musics"
	"middleware/example/internal/helpers"
	_ "middleware/example/internal/models"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

func main() {
	r := chi.NewRouter()

	r.Route("/musics", func(r chi.Router) {
		r.Get("/", musics.GetMusics)
		r.Post("/", musics.PostMusic)
		r.Route("/{id}", func(r chi.Router) {
			r.Use(musics.Ctx)
			r.Get("/", musics.GetMusic)
			r.Delete("/", musics.DeleteMusic)
			r.Put("/", musics.PutMusic)

		})
	})

	r.Route("/artists", func(r chi.Router) {
		r.Get("/", artists.GetArtists)
		r.Post("/", artists.PostArtist)
		r.Route("/{id}", func(r chi.Router) {
			r.Use(artists.Ctx)
			r.Get("/", artists.GetArtist)
			r.Delete("/", artists.DeleteArtist)
			r.Put("/", artists.PutArtist)
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
			id CHAR(36) PRIMARY KEY,
			name TEXT NOT NULL
		);
		
		CREATE TABLE IF NOT EXISTS Artist (
			id CHAR(36) PRIMARY KEY,
			name TEXT NOT NULL
		);
		
		CREATE TABLE IF NOT EXISTS Album (
			id CHAR(36) PRIMARY KEY,
			name TEXT NOT NULL,
			artistId CHAR(36),
			FOREIGN KEY (artistId) REFERENCES Artist(id)
		);
		
		CREATE TABLE IF NOT EXISTS Music (
			id CHAR(36) PRIMARY KEY,
			title TEXT NOT NULL,
			genreId CHAR(36),
			artistId CHAR(36),
			albumId CHAR(36),
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
