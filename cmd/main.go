package main

import (
	"fmt"
	"middleware/example/internal/controllers/users"
	"middleware/example/internal/helpers"
	_ "middleware/example/internal/models"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid"
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
	// TODO: delete that part when doing Flask
	dropStatements := []string{
		"DROP TABLE IF EXISTS users;",
	}

	for _, stmt := range dropStatements {
		if _, err := db.Exec(stmt); err != nil {
			logrus.Fatalf("Could not drop table! Error was: %s", err.Error())
		}
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

	for i := 0; i < 2; i++ {
		uuid, err := uuid.NewV4()
		if err != nil {
			logrus.Fatalf("failed to generate UUID: %s", err.Error())
		}

		username := fmt.Sprintf("testuser%d", i+1)
		email := fmt.Sprintf("test%d@example.com", i+1)

		insertUserSQL := `INSERT INTO users (id, username, email) VALUES (?, ?, ?)`
		if _, err := db.Exec(insertUserSQL, uuid.String(), username, email); err != nil {
			logrus.Errorf("could not insert test user %s: %s", username, err.Error())
		}
	}
	helpers.CloseDB(db)
}
