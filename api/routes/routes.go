package routes

import (
	"github.com/gorilla/mux"

	"github.com/victor-nach/nigerian-senators/api/controllers"
)

func Router() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	// router.HandleFunc("/", controllers.Welcome).Methods("GET")
	router.HandleFunc("/", controllers.Welcome).Methods("GET")
	router.HandleFunc("/senators", controllers.GetAllSenators).Methods("GET")

	return router
}