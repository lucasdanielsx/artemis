package api

import (
	util "artemis/pkg/util"
	"net/http"
)

func Health(w http.ResponseWriter, request *http.Request) {
	w.WriteHeader(200)
	_, err := w.Write([]byte("OK"))

	util.HandleError("Internal Error", err)
}
