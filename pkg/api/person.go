package api

import "net/http"

func GetPersonServer(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func DeletePersonServer(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}
