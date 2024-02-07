package todo

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type handler struct {
	service InterfaceService
}

func NewHandler(sv InterfaceService) handler {

	return handler{sv}
}

func (h *handler) GetTodoByID(c echo.Context) error {
	response, err := h.service.GetTodoByID(1)
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, *response)
}

func (h *handler) GetTodoList(c echo.Context) error {
	response, err := h.service.GetTodoList()
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, response)
}
