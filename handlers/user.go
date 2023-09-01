package handlers

import (
	"net/http"

	"github.com/Cristi571/golang_api/services"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}

/*REMOVE THIS ENDPOINT*/
func (h *UserHandler) Get(ctx echo.Context) error {
	id := ctx.Param("id")

	user, err := h.userService.GetUser(ctx.Request().Context(), id)
	if err != nil {
		ctx.NoContent(http.StatusInternalServerError)
		return nil
	}

	if user == nil {
		ctx.JSON(http.StatusNotFound, map[string]interface{}{"message": "user not found"})
		return nil
	}

	ctx.JSON(http.StatusOK, user)
	return nil
}

/*IMPLEMENT THIS ENDPOINT*/
func (h *UserHandler) Post(ctx echo.Context) error {
	// id := ctx.Param("id")

	// user, err := h.userService.GetUser(ctx.Request().Context(), id)

	return nil
}
