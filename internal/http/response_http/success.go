package responsehttp

import (
	"net/http"

	"github.com/MnPutrav2/go_architecture/pkg/response"
)

func SuccessCreated(message string, w http.ResponseWriter, r *http.Request) {
	response.Message(message, "success", "INFO", http.StatusCreated, w, r)
}
