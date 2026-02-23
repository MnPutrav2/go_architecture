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
			middleware.RateLimiter(1, 1, middleware.Authorization(w, r, h.Create))
		default:
			response.Message("method not allowed", "method not allowed", "WARN", http.StatusMethodNotAllowed, w, r)
		}
	}
}
