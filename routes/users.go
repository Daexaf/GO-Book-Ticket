package routes

import (
	"net/http"
	"strconv"

	"example.com/BookEvent/utils"

	"example.com/BookEvent/models"

	"github.com/gin-gonic/gin"
)

func signup(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error message": "Bad request"})
		return
	}

	err = user.Save()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error message": "Cant save user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func login(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Can't authenticate user"})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.Id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error message": "Cant create token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User logged in successfully", "token": token})
}

func getAllUsers(c *gin.Context) {
	users, err := models.GetAllUsers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error message": "Cant fetch users"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}

func getUserById(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}

	user, err := models.GetUserById(userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error message": "Cant fetch user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func deleteUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}

	user, err := models.GetUserById(userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error message": "Cant fetch user"})
		return
	}

	err = user.Deletes()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error message": "Cant delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
