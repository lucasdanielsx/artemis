package api

import (
	"artemis/pkg/database"
	"artemis/pkg/util"
	"encoding/json"
	"net/http"
)

type health struct {
	Name   string `json:"name"`
	Status bool   `json:"status"`
}

func HealthServer(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var checks []health

	redis := database.RedisHealth()
	postgres := database.PostgresHealth()

	checks = append(checks, health{"redis", redis})
	checks = append(checks, health{"postgres", postgres})

	if redis && postgres {
		checks = append(checks, health{"app", true})
		w.WriteHeader(200)
	} else {
		checks = append(checks, health{"app", false})
		w.WriteHeader(503)
	}

	marshal, err := json.Marshal(checks)
	util.HandleError("Can not marshal checks", err)

	_, err = w.Write(marshal)
	util.HandleError("Can not encode checks", err)
}
