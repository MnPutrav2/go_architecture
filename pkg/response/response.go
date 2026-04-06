package response

import (
	"net/http"

	modelresponse "github.com/MnPutrav2/go_architecture/pkg/response/model_response"
)

func Created(message string, w http.ResponseWriter, r *http.Request) {
	modelresponse.Message(message, "Success", "INFO", http.StatusCreated, w, r)
}

func BadRequest(message string, err error, w http.ResponseWriter, r *http.Request) {
	modelresponse.Message(message, err.Error(), "WARN", http.StatusBadRequest, w, r)
}

func Unauthorization(message string, err error, w http.ResponseWriter, r *http.Request) {
	modelresponse.Message(message, err.Error(), "WARN", http.StatusUnauthorized, w, r)
}

func ToManyRequest(message string, err error, w http.ResponseWriter, r *http.Request) {
	modelresponse.Message(message, err.Error(), "WARN", http.StatusTooManyRequests, w, r)
}

func Forbidden(message string, err error, w http.ResponseWriter, r *http.Request) {
	modelresponse.Message(message, err.Error(), "WARN", http.StatusForbidden, w, r)
}

// Add more
