package api

import (
	"log"
	"net/http"
	"os"

	"github.com/victor-nach/nigerian-senators/api/routes"
)


func Run() {
	PORT, ok := os.LookupEnv("PORT")
	if !ok {
		PORT = "8080"
	}
	log.Println("serving on port:", PORT)

	log.Fatal(http.ListenAndServe(":"+PORT, routes.Router()))
}