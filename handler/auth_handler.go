package handler

import (
	"github.com/MuhammadIbraAlfathar/gin-api/dto"
	"github.com/MuhammadIbraAlfathar/gin-api/errorhandler"
	"github.com/MuhammadIbraAlfathar/gin-api/helper"
	"github.com/MuhammadIbraAlfathar/gin-api/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type authHandler struct {
	service service.AuthService
}

func NewAuthHandler(s service.AuthService) *authHandler {
	return &authHandler{
		service: s,
	}
}

func (h *authHandler) Register(c *gin.Context) {
	var register dto.RegisterRequest

	if err := c.ShouldBindJSON(&register); err != nil {
		errorhandler.HandlerError(c, &errorhandler.BadRequestError{
			Message: err.Error(),
		})
		return
	}

	if err := h.service.Register(&register); err != nil {
		errorhandler.HandlerError(c, err)
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "Register Successfully",
	})

	c.JSON(http.StatusCreated, res)

}
