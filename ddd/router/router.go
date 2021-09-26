package router

import (
	"database/sql"

	"github.com/alochym01/thecodecamp_1/domain/users"
	// "github.com/alochym01/thecodecamp_1/storage/memory"
	"github.com/alochym01/thecodecamp_1/storage/sqlite"
	"github.com/gin-gonic/gin"
)

// Router ...
func Router(db *sql.DB) *gin.Engine {
	router := gin.Default()

	// uRepo := memory.NewRepository()
	uRepo := sqlite.NewRepository(db)
	uSVCRepo := users.NewService(uRepo)
	uHandler := users.NewHandler(uSVCRepo)
	router.GET("/users", uHandler.GetUsers)
	router.GET("/users/:email", uHandler.GetUserByEmail)
	// router.GET("/users/:id", uHandler.GetUserByID)

	return router
}
