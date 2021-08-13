package main

import (
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
)

func main() {
	var Error = log.New(os.Stdout, "\u001b[31mERROR: \u001b[0m", log.LstdFlags|log.Lshortfile)
	var Info = log.New(os.Stdout, "\u001b[34mINFO: \u001B[0m", log.LstdFlags|log.Lshortfile)

	viper.SetConfigFile("./config/.env")
	dotEnvErr := viper.ReadInConfig()

	if dotEnvErr != nil {
		Error.Fatal("Error loading ..env file : ", dotEnvErr)
	}

	port, _ := viper.Get("ARTEMIS_PORT").(string)

	Info.Println("Artemis API")

	router := mux.NewRouter()
	err := http.ListenAndServe(port, router)

	if err != nil {
		Error.Fatal("Can not start Artemis API: ", err)
	}
}
