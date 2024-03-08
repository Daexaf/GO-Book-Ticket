package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/BookEvent/models"

	"github.com/gin-gonic/gin"
)

func registerEvent(c *gin.Context) {
	userId := c.GetInt64("userId")
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "Could not parse event Id."})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not fetch events."})
		return
	}

	if !isUserAuthorizedToRegister(c, userId, event.UserId) {
		c.JSON(http.StatusUnauthorized, gin.H{"Message": "Not Authorized"})
		return
	}

	err = event.Register(userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Message": "could not register user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Registered"})
}

func cancleRegisterEvent(c *gin.Context) {
	userId := c.GetInt64("userId")
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "Could not parse event Id."})
		return
	}

	var event models.Event
	event.Id = eventId

	event.CancleRegisterEvent(userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Message": "could not cancel register user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Canceled"})
}

func isUserAuthorizedToRegister(c *gin.Context, userId, eventUserId int64) bool {
	// Tambahkan logika otorisasi sesuai kebutuhan Anda
	// Contoh sederhana: Izinkan pengguna mendaftar ke acara jika userId dan eventUserId sama
	fmt.Println("userId:", userId)
	fmt.Println("eventUserId:", eventUserId)
	return userId == eventUserId
}
