package main

import (
	"artemis/pkg/api"
	"artemis/pkg/util"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"net/http"
)

func handleRequests(port string) {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/health", api.HealthServer)
	myRouter.HandleFunc("/token", api.TokenizeServer).Methods("POST")
	myRouter.HandleFunc("/token/{value}", api.PersistTokenServer).Methods("POST")
	myRouter.HandleFunc("/token/{value}", api.DeleteTokenServer).Methods("DELETE")
	myRouter.HandleFunc("/person/{value}", api.GetPersonServer)
	myRouter.HandleFunc("/person/{value}", api.DeletePersonServer).Methods("DELETE")

	err := http.ListenAndServe(port, myRouter)
	util.HandleFatal("Can not start Artemis API: ", err)
}

func main() {
	viper.SetConfigFile("./config/.env")

	err := viper.ReadInConfig()
	util.HandleFatal("Error loading .env file : ", err)

	port, _ := viper.Get("ARTEMIS_PORT").(string)

	util.Info("Starting Artemis API on port " + port)

	handleRequests(port)
}
