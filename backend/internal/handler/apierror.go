package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type APIErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func apiError(w http.ResponseWriter, status int, msg string, a ...any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(APIErrorResponse{
		Status:  status,
		Message: fmt.Sprintf(msg, a...),
	})
}
