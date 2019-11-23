package models

import (
	
)

type Response struct {
	Status  int         `json:"status,omitempty"`
	Message string      `json:"message,omitempty"`
	Error	string		`json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}