package api

import (
	database "artemis/pkg/database"
	"artemis/pkg/util"
	"encoding/json"
	"net/http"
)

type Health struct {
	Name   string `json:"name"`
	Status bool   `json:"status"`
}

func HealthServer(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var checks []Health

	redis := database.RedisHealth()
	postgres := database.PostgresHealth()

	checks = append(checks, Health{"redis", redis})
	checks = append(checks, Health{"postgres", postgres})

	if !redis || !postgres {
		checks = append(checks, Health{"app", false})
		w.WriteHeader(503)
	} else {
		checks = append(checks, Health{"app", true})
		w.WriteHeader(200)
	}

	marshal, err := json.Marshal(checks)
	util.HandleError("Can not marshal checks", err)

	_, err = w.Write(marshal)
	util.HandleError("Can not encode checks", err)
}
