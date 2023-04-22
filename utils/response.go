package utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string      `json:"message"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data,omitempty"`
}

// FOR TESTING PURPOSE
type ResponseMap struct {
	Message string                 `json:"message"`
	Status  string                 `json:"status"`
	Data    map[string]interface{} `json:"data,omitempty"`
}

// FOR TESTING PURPOSE
type ResponseSlice struct {
	Message string                   `json:"message"`
	Status  string                   `json:"status"`
	Data    []map[string]interface{} `json:"data,omitempty"`
}

var HttpMessage = map[int]string{
	200: "Success",
	201: "Success",
	202: "Accepted",
	204: "No Content",
	400: "Bad Request",
	401: "Unauthorized",
	403: "Forbidden",
	404: "Not Found",
	500: "Internal Server Error",
}

func GenerateJsonResponse(w http.ResponseWriter, data interface{}, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := Response{
		Message: message,
		Status:  HttpMessage[statusCode],
		Data:    data,
	}

	responseEncode, err := json.Marshal(response)
	PanicIfAppError(err, "failed when marshar response", 500)
	w.Write(responseEncode)
}
