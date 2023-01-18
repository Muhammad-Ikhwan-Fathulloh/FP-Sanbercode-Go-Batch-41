package router

import (
	"FP-Sanbercode-Go-Batch-41/controller"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/token", controller.GenerateToken)
	// Route Users
	router.GET("/users", controller.GetAllUser)
	router.POST("/users", controller.InsertUser)
	router.PUT("/users/:id", controller.UpdateUser)
	router.DELETE("/users/:id", controller.DeleteUser)

	// Route Communities
	router.GET("/communities", controller.GetAllCommunity)
	router.POST("/communities", controller.InsertCommunity)
	router.PUT("/communities/:id", controller.UpdateCommunity)
	router.DELETE("/communities/:id", controller.DeleteCommunity)

	// Route Event Categories
	router.GET("/event-categories", controller.GetAllEventCategory)
	router.POST("/event-categories", controller.InsertEventCategory)
	router.PUT("/event-categories/:id", controller.UpdateEventCategory)
	router.DELETE("/event-categories/:id", controller.DeleteEventCategory)

	// Route Events
	router.GET("/events", controller.GetAllEvent)
	router.POST("/events", controller.InsertEvent)
	router.PUT("/events/:id", controller.UpdateEvent)
	router.DELETE("/events/:id", controller.DeleteEvent)

	// Route Event Participants
	router.GET("/event-participants", controller.GetAllEventParticipant)
	router.POST("/event-participants", controller.InsertEventParticipant)
	router.PUT("/event-participants/:id", controller.UpdateEventParticipant)
	router.DELETE("/event-participants/:id", controller.DeleteEventParticipant)

	return router
}
