package main

import (
	api "artemis/pkg/api"
	util "artemis/pkg/util"
	"github.com/spf13/viper"
	"net/http"
)

func main() {
	viper.SetConfigFile("./config/.env")

	err := viper.ReadInConfig()
	util.HandleFatal("Error loading .env file : ", err)

	port, _ := viper.Get("ARTEMIS_PORT").(string)

	util.Info("Starting Artemis API on port " + port)

	http.HandleFunc("/health", api.HealthServer)

	err = http.ListenAndServe(port, nil)
	util.HandleFatal("Can not start Artemis API: ", err)
}
