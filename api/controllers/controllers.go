package controllers

import (
	"net/http"
	"encoding/json"

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