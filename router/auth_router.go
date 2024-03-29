package router

import (
	"github.com/MuhammadIbraAlfathar/gin-api/config"
	"github.com/MuhammadIbraAlfathar/gin-api/handler"
	"github.com/MuhammadIbraAlfathar/gin-api/repository"
	"github.com/MuhammadIbraAlfathar/gin-api/service"
	"github.com/gin-gonic/gin"
)

func AuthRouter(api *gin.RouterGroup) {
	authRepository := repository.NewAuthRepository(config.DB)
	authService := service.NewAuthService(authRepository)
	authHandler := handler.NewAuthHandler(authService)

	api.POST("/register", authHandler.Register)

}
