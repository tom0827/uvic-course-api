package utils

import (
	"course-api/models"
	"encoding/json"
	"net/http"
)

func WriteSuccess(w http.ResponseWriter, data interface{}) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.APIResponse{
		Status: "success",
		Data:   data,
	})
}

func WriteError(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(models.APIResponse{
		Status:  "error",
		Message: message,
		Data:    nil,
	})
}
