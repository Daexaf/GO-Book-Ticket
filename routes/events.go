package routes

import (
	"net/http"
	"strconv"

	"example.com/BookEvent/models"
	"github.com/gin-gonic/gin"
)

func getAllEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not fetch events, try again later."})
		return
	}
	c.JSON(http.StatusOK, events)
}

func getEventById(c *gin.Context) {
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
	c.JSON(http.StatusOK, event)
}

func createEvent(c *gin.Context) {

	var event models.Event
	err := c.ShouldBindJSON(&event)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error message": err.Error()})
		return
	}

	userId := c.GetInt64("userId")
	event.UserId = userId

	err = event.Save()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not fetch events, try again later."})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Event Created", "event": event})
}

func updateEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "Could not parse event Id."})
		return
	}

	userId := c.GetInt64("userId")
	event, err := models.GetEventById(eventId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not fetch events."})
		return
	}

	if event.UserId != userId {
		c.JSON(http.StatusUnauthorized, gin.H{"Message": "Not authorized to update event."})
		return
	}

	var updatedEvent models.Event
	err = c.ShouldBindJSON(&updatedEvent)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error message": err.Error()})
		return
	}

	updatedEvent.Id = eventId
	updatedEvent.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not update events, try again later."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "Updated events Successfully"})
}

func deleteEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "Could not parse event Id."})
		return
	}

	userId := c.GetInt64("userId")
	event, err := models.GetEventById(eventId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not fetch events."})
		return
	}

	if event.UserId != userId {
		c.JSON(http.StatusUnauthorized, gin.H{"Message": "Not authorized to delete event."})
		return
	}

	err = event.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not delete events, try again later."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "Event deleted successfully"})
}
