package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/MnPutrav2/go_architecture/internal/model"
	"github.com/MnPutrav2/go_architecture/internal/service"
	"github.com/MnPutrav2/go_architecture/pkg/decoder"
	"github.com/MnPutrav2/go_architecture/pkg/response"
)

func CreateUserHandler(service service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, close := context.WithTimeout(r.Context(), time.Second*5)
		defer close()

		body, err := decoder.BodyDecoder[model.CreateUser](r)
		if err != nil {
			response.BadRequest("Failed decode body", err, w, r)
			return
		}

		if err := service.CreateUserService(ctx, body); err != nil {
			response.BadRequest("Failed create account", err, w, r)
			return
		}

		response.Created("Success", w, r)
	}
}

// Entry
