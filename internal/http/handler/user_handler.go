package handler

import (
	"context"
	"net/http"

	"github.com/MnPutrav2/go_architecture/internal/helper"
	"github.com/MnPutrav2/go_architecture/internal/model"
	"github.com/MnPutrav2/go_architecture/internal/service"
	"github.com/MnPutrav2/go_architecture/pkg/prefix"
	res "github.com/MnPutrav2/go_architecture/pkg/response"
	"github.com/MnPutrav2/go_architecture/pkg/validator"
)

// Entry

func GetUserHandler(service service.UserService) http.HandlerFunc {
	return helper.Handler(func(ctx context.Context, w http.ResponseWriter, r *http.Request) {

		result, err := service.GetUserService(ctx)
		if err != nil {
			res.BadRequest("Failed create account", err, w, r)
			return
		}

		res.Data("Success", result, w, r)

	})
}

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

func DeleteUserHandler(service service.UserService) http.HandlerFunc {
	return helper.Handler(func(ctx context.Context, w http.ResponseWriter, r *http.Request) {

		param, err := prefix.UUID("id", r)
		if err != nil {
			res.BadRequest(err.Error(), err, w, r)
			return
		}

		if err := service.DeleteUserService(ctx, param); err != nil {
			res.BadRequest("Failed remove account", err, w, r)
			return
		}

		res.Success("Success", w, r)

	})
}
