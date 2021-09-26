package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserHandler ...
type UserHandler struct {
	us UserServiceRepo
}

// GetUsers ...
func (u UserHandler) GetUsers(c *gin.Context) {
	allusers, err := u.us.GetUsers()
	if err != nil {
		// if err == sql.ErrNoRows {
		// 	fmt.Println("Not Found", err.Error())
		// 	c.JSON(http.StatusNotFound, gin.H{
		// 		"err": sql.ErrNoRows.Error(),
		// 	})
		// 	return
		// }
		// fmt.Println("Server Err", err.Error())
		// c.JSON(http.StatusInternalServerError, gin.H{
		// 	"err": err.Error(),
		// })
		c.JSON(err.Code, gin.H{
			"err": err.Message,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": allusers,
	})
}

// GetUserByEmail ...
func (u UserHandler) GetUserByEmail(c *gin.Context) {
	email := c.Param("email")
	user, err := u.us.GetUserByEmail(email)

	if err != nil {
		// if err == sql.ErrNoRows {
		// 	fmt.Println("Not Found", err.Error())
		// 	c.JSON(http.StatusNotFound, gin.H{
		// 		"err": sql.ErrNoRows.Error(),
		// 	})
		// 	return
		// }
		// fmt.Println("Server Err", err.Error())
		// c.JSON(http.StatusInternalServerError, gin.H{
		// 	"err": err.Error(),
		// })
		c.JSON(err.Code, gin.H{
			"err": err.Message,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

// NewHandler return a UserServiceRepo
func NewHandler(uRepo UserServiceRepo) *UserHandler {
	return &UserHandler{
		us: uRepo,
	}
}
