package api

import (
	"log"
	"net/http"

	"github.com/victor-nach/nigerian-senators/api/routes"
)


func Run() {
	log.Println("welcome")
	PORT := "8080"
	log.Println("serving on port:", PORT)

	log.Fatal(http.ListenAndServe(":"+PORT, routes.Router()))
}