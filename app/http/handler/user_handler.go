package handler

import (
	"context"
	"net/http"

	"github.com/MnPutrav2/go_architecture/app/helper"
	"github.com/MnPutrav2/go_architecture/app/model"
	res "github.com/MnPutrav2/go_architecture/app/pkg/response"
	"github.com/MnPutrav2/go_architecture/app/pkg/validator"
	"github.com/MnPutrav2/go_architecture/app/service"
)

// Entry

func CreateUserHandler(service service.UserService) http.HandlerFunc {
	return helper.Handler(func(ctx context.Context, w http.ResponseWriter, r *http.Request) {

		body, err := validator.ValidatePayload[model.CreateUser](r)
		if err != nil {
			res.BadRequest(err.Error(), err, w, r)
			return
		}

		if err := service.CreateUserService(ctx, body); err != nil {
			res.BadRequest("Failed create account", err, w, r)
			return
		}

		res.Created("Success", w, r)

	})
}
