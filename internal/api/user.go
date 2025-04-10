package api

import (
	"encoding/json"
	"net/http"
)

// User represents a simple user model
type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// GetUsers returns a JSON response with user data
func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := []User{
		{ID: "1", Name: "Alice"},
		{ID: "2", Name: "Boba"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
