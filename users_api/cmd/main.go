package main

import (
	"middleware/example2/internal/controllers/users"
	"middleware/example2/internal/helpers"
	_ "middleware/example2/internal/models"
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "api/swagger.json")
	})

	r.Route("/users", func(r chi.Router) {
		r.Get("/", users.GetUsers)
		r.Route("/{id}", func(r chi.Router) {
			r.Use(users.Ctx)
			r.Get("/", users.GetUserById)
		})
		r.Post("/", users.CreateUser)
		r.Put("/{id}", users.UpdateUser)
		r.Delete("/{id}", users.DeleteUser)
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
		`CREATE TABLE IF NOT EXISTS users (
		  	id UUID PRIMARY KEY,
  			username VARCHAR(255) NOT NULL UNIQUE,
  			email VARCHAR(255) NOT NULL UNIQUE,
  			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);`,
	}
	for _, scheme := range schemes {
		if _, err := db.Exec(scheme); err != nil {
			logrus.Fatalln("Could not generate table ! Error was : " + err.Error())
		}
	}
	helpers.CloseDB(db)
}
