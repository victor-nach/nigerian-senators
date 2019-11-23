package controllers

import (
	"net/http"
	"encoding/json"
	"log"

	"github.com/victor-nach/nigerian-senators/api/models"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	response := models.Response{
		Status:  200,
		Message: "Welcome !",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetAllSenators(w http.ResponseWriter, r *http.Request) {
	Senators, err := models.Senator.GetAll(models.Senator{})
	if err != nil {
		log.Println(err)
	}
	log.Println(Senators)

	response := models.Response{
		Status:  200,
		Message: "All senators",
		Data: Senators,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}