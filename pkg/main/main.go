package main

import (
	Logs "artemis/pkg/util"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"net/http"
)

func main() {
	viper.SetConfigFile("./config/.env")
	dotEnvErr := viper.ReadInConfig()

	if dotEnvErr != nil {
		Logs.Fatal("Error loading ..env file : ", dotEnvErr)
	}

	port, _ := viper.Get("ARTEMIS_PORT").(string)

	Logs.Info("Artemis API")

	router := mux.NewRouter()
	err := http.ListenAndServe(port, router)

	if err != nil {
		Logs.Fatal("Can not start Artemis API: ", err)
	}
}
