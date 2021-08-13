package main

import (
	api "artemis/pkg/api"
	util "artemis/pkg/util"
	"github.com/spf13/viper"
	"net/http"
)

func main() {
	util.Info("Artemis API")
	viper.SetConfigFile("./config/.env")

	err := viper.ReadInConfig()
	util.HandleFatal("Error loading .env file : ", err)

	port, _ := viper.Get("ARTEMIS_PORT").(string)

	http.HandleFunc("/health", api.Health)

	err = http.ListenAndServe(port, nil)
	util.HandleFatal("Can not start Artemis API: ", err)
}
