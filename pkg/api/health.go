package api

import (
	database "artemis/pkg/database"
	"net/http"
)

func Health(w http.ResponseWriter, _ *http.Request) {
	redis := database.RedisHealth()
	postgres := database.PostgresHealth()

	if redis && postgres {
		w.WriteHeader(200)
	}

	w.WriteHeader(500)
}
