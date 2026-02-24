package route

import (
	"net/http"

	"github.com/MnPutrav2/go_architecture/internal/http/handler"
	userService "github.com/MnPutrav2/go_architecture/internal/service/user"
	"github.com/MnPutrav2/go_architecture/pkg/middleware"
	"github.com/MnPutrav2/go_architecture/pkg/response"
)

func UserRoute(service userService.UserService) http.HandlerFunc {
	h := handler.InitUserHandle(service)

	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			middleware.Chain(h.Create,
				middleware.CTJson,
				middleware.Authorization,
				middleware.CORS,
			)(w, r)
		default:
			response.Message("method not allowed", "method not allowed", "WARN", http.StatusMethodNotAllowed, w, r)
		}
	}
}
