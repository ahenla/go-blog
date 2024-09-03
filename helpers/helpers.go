package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ParseJSON(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}

	return json.NewDecoder(r.Body).Decode(payload)
}

func RespondJSON(w http.ResponseWriter, status int, content any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(content)
}

func RespondError(w http.ResponseWriter, status int, content error) {
	RespondJSON(w, status, map[string]string{"error": content.Error()})
}
