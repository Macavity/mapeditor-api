package Users

import (
	"Macavity/mapeditor-server/server/Users/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func RegisterUserRoutes(route *gin.Engine) {
	route.GET("/users/", getAllUsersRoute)
	route.GET("/users/:id", getUserByIdRoute)
	route.POST("/users", createUserRoute)
}

func getAllUsersRoute(c *gin.Context) {
	users, err := FindAllUsers()
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

func createUserRoute(c *gin.Context) {
	var input models.CreateUserDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	user := models.User{
		Name:  input.Name,
		Email: input.Email,
	}
	CreateUser(&user)

	c.JSON(http.StatusOK, gin.H{"user": &user})
}

func getUserByIdRoute(c *gin.Context) {
	userId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	user, err := FindUserById(userId)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": &user})
}
