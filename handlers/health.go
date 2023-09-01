package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) IsAlive(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{"message": "Okkkkkkk"})
}



func HomePage(w http.ResponseWriter, r *http.Request){
	// return ctx.JSON(http.StatusOK, map[string]string{"message": "Welcome to the HomePage!"})
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}