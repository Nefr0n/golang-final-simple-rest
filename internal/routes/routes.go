package routes

import (
	"github.com/Skapar/simple-rest/internal/api"
	"github.com/Skapar/simple-rest/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// SetupRoutes - настройка маршрутов
func SetupRoutes(router *gin.Engine, authService service.AuthService, userService service.UserService) {
	healthHandler := api.NewHealthHandler()

	authHandler := api.NewAuthHandler(authService)
	userHandler := api.NewUserHandler(userService)

	apiV1 := router.Group("/api/v1")

	// routes
	// health
	router.GET("/health", healthHandler.CheckHealthHandler)

	// user
	apiV1.POST("/signup", authHandler.RegisterUserHandler)
	apiV1.GET("/user/:user_id", userHandler.GetUserHandler)
	apiV1.PUT("/user/:user_id", userHandler.UpdateUserHandler)
	apiV1.DELETE("/user/:user_id", userHandler.DeleteUserHandler)
	apiV1.PATCH("/user/:user_id", userHandler.SoftDeleteUserHandler)

	// Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
