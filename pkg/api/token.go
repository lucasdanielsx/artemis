package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func TokenizeServer(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func PersistTokenServer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	value := vars["value"]
	fmt.Fprintf(w, "Key: "+value)
	w.Header().Set("Content-Type", "application/json")
}

func DeleteTokenServer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	value := vars["value"]
	fmt.Fprintf(w, "Key: "+value)
	w.Header().Set("Content-Type", "application/json")
}
