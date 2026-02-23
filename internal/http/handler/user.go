package handler

import (
	"context"
	"net/http"
	"time"

	userModel "github.com/MnPutrav2/go_architecture/internal/model/user"
	userService "github.com/MnPutrav2/go_architecture/internal/service/user"
	"github.com/MnPutrav2/go_architecture/pkg/response"
	"github.com/MnPutrav2/go_architecture/pkg/util"
)

type UserHandle struct {
	service userService.UserService
}

func InitUserHandle(service userService.UserService) *UserHandle {
	return &UserHandle{service: service}
}

// Entry

func (h *UserHandle) Create(w http.ResponseWriter, r *http.Request) {

	ctx, close := context.WithTimeout(r.Context(), time.Second*5)
	defer close()

	payload, err := util.BodyDecoder[userModel.Create](r)
	if err != nil {
		response.Message("failed decode body", err.Error(), "WARN", http.StatusMethodNotAllowed, w, r)
		return
	}

	if err := h.service.CreateUser(ctx, payload); err != nil {
		response.Message("failed create user", err.Error(), "WARN", http.StatusMethodNotAllowed, w, r)
		return
	}

	response.Message("success", "success", "INFO", http.StatusCreated, w, r)
}
