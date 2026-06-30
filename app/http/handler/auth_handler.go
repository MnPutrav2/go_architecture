package handler

import (
	"context"
	"errors"
	"net/http"

	"github.com/MnPutrav2/go_architecture/app/helper"
	httperror "github.com/MnPutrav2/go_architecture/app/http/error"
	"github.com/MnPutrav2/go_architecture/app/model"
	res "github.com/MnPutrav2/go_architecture/app/pkg/response"
	"github.com/MnPutrav2/go_architecture/app/pkg/validator"

	"github.com/MnPutrav2/go_architecture/app/service"
)

func LoginUserHandler(service service.AuthService) http.HandlerFunc {
	return helper.Handler(func(ctx context.Context, w http.ResponseWriter, r *http.Request) {

		body, err := validator.ValidatePayload[model.LoginUser](r)
		if err != nil {
			res.BadRequest(err.Error(), err, w, r)
			return
		}

		result, err := service.LoginService(ctx, body)
		if err != nil {
			res.BadRequest("Login failed", err, w, r)
			return
		}

		res.Data("success", result, w, r)

	})
}

func RefreshTokenHandler(service service.AuthService) http.HandlerFunc {
	return helper.Handler(func(ctx context.Context, w http.ResponseWriter, r *http.Request) {

		body, err := validator.ValidatePayload[model.RefreshTokenRequest](r)
		if err != nil {
			res.BadRequest(err.Error(), err, w, r)
			return
		}

		result, err := service.RefreshToken(ctx, body.RefreshToken)
		if err != nil {
			if errors.Is(err, httperror.ErrIsUnauthorization) {
				res.Unauthorization("unauthorization", err, w, r)
				return
			}

			res.BadRequest("Failed", err, w, r)
			return
		}

		res.Data("success", result, w, r)

	})
}

// Entry
