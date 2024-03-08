package routes

import (
	"example.com/BookEvent/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	//routes event
	server.GET("/events", getAllEvents)
	server.GET("/events/:id", getEventById)

	//login & signup
	server.POST("/signup", signup)
	server.POST("/login", login)

	//routes user
	server.GET("/users", getAllUsers)
	server.GET("/users/:id", getUserById)

	//protected routes
	authenticated := server.Group("/")
	authenticated.Use(middleware.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerEvent)
	authenticated.DELETE("/events/:id/register", cancleRegisterEvent)
	authenticated.DELETE("/users/:id", deleteUser)
}
