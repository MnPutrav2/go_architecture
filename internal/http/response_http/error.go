package responsehttp

import (
	"net/http"

	"github.com/MnPutrav2/go_architecture/pkg/response"
)

func ErrorBadRequest(message string, err error, w http.ResponseWriter, r *http.Request) {
	if err != nil {
		response.Message(message, err.Error(), "WARN", http.StatusBadRequest, w, r)
		return
	}
}
