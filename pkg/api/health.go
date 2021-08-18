package api

import (
	database "artemis/pkg/database"
	"net/http"
)

func Health(w http.ResponseWriter, request *http.Request) {
	_, err := database.RedisHealth()

	if err != nil {
		w.WriteHeader(200)
	}

	w.WriteHeader(500)
}
