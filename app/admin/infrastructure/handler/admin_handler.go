package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type AdminHandler struct {
}

func NewAdminHandler() *AdminHandler {
	return &AdminHandler{}
}

func (h *AdminHandler) Ping(ectx echo.Context) error {
	return ectx.NoContent(http.StatusOK)
}
