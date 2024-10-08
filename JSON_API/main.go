package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	ID    int
	Valid bool
}
type apiError struct {
	err    string `json:"error"`
	status int
}

type apiFunc func(http.ResponseWriter, *http.Request) error

func main() {

	http.HandleFunc("/user", HttpHandleMidlleware(handleGetUserByID))

	http.ListenAndServe(":3000", nil)
}

func HttpHandleMidlleware(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJson(w, http.StatusInternalServerError, apiError{err: "Internal Server Error"})
			return
		}
	}
}

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

func handleGetUserByID(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodGet {
		return WriteJson(w, http.StatusOK, apiError{err: "Method Not Allowed"})
	}

	// db
	user := User{
		ID:    1,
		Valid: false,
	}
	if !user.Valid {
		return WriteJson(w, http.StatusForbidden, apiError{err: "forbidden"})
	}

	return WriteJson(w, http.StatusOK, user)
}
