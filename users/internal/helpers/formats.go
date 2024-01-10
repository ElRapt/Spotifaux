package helpers

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"gopkg.in/yaml.v2"
)

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
