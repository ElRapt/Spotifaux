package helpers

import (
	"database/sql"
	"encoding/json"
	"encoding/xml"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

func OpenDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "file:ratings.db")
	if err != nil {
		db.SetMaxOpenConns(1)
	}
	return db, err
}
func CloseDB(db *sql.DB) {
	err := db.Close()
	if err != nil {
		logrus.Errorf("error closing db : %s", err.Error())
	}
}

func RespondWithFormat(w http.ResponseWriter, r *http.Request, data interface{}) {
	switch r.Header.Get("Accept") {
	case "application/xml":
		w.Header().Set("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(data)
	case "application/yaml":
		w.Header().Set("Content-Type", "application/yaml")
		yaml.NewEncoder(w).Encode(data)

	default:
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	}
}
