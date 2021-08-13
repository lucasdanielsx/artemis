package main

import (
	Logs "artemis/pkg/util"
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

	err := http.ListenAndServe(port, nil)

	if err != nil {
		Logs.Fatal("Can not start Artemis API: ", err)
	}
}
