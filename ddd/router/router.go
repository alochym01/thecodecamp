package router

import (
	"github.com/alochym01/thecodecamp_1/domain/users"
	"github.com/alochym01/thecodecamp_1/storage/memory"
	"github.com/gin-gonic/gin"
)

// Router ...
func Router() *gin.Engine {
	router := gin.Default()

	uRepo := memory.NewRepository()
	// uSVCRepo := service.NewUserService(uRepo)
	// uHandler := handlers.NewUserHandler(uSVCRepo)
	uSVCRepo := users.NewUserService(uRepo)
	uHandler := users.NewUserHandler(uSVCRepo)
	router.GET("/users", uHandler.GetUsers)
	router.GET("/users/:email", uHandler.GetUserByEmail)

	return router
}
