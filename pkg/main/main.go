package main

import (
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	var Error = log.New(os.Stdout, "\u001b[31mERROR: \u001b[0m", log.LstdFlags|log.Lshortfile)
	var Info = log.New(os.Stdout, "\u001b[34mINFO: \u001B[0m", log.LstdFlags|log.Lshortfile)

	dotEnvErr := godotenv.Load("./config/..env")

	if dotEnvErr != nil {
		Error.Fatal("Error loading ..env file : ", dotEnvErr)
	}

	Info.Println("Artemis API")

	router := mux.NewRouter()
	err := http.ListenAndServe(os.Getenv("ARTEMIS_PORT"), router)

	if err != nil {
		Error.Fatal("Can not start Artemis API: ", err)
	}
}
