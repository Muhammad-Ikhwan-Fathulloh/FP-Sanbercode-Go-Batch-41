package router

import (
	"FP-Sanbercode-Go-Batch-41/controller"
	"FP-Sanbercode-Go-Batch-41/middleware"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		// Route Login Public
		api.POST("/login", controller.GenerateToken)

		// Route Community Public
		api.GET("/communities", controller.GetAllCommunity)
		api.GET("/communities/:community_id", controller.GetCommunityById)

		// Route Event Category Public
		api.GET("/event-categories", controller.GetAllEventCategory)
		api.GET("/event-categories/:event_category_id", controller.GeteventCategoryById)

		// Route Event Public
		api.GET("/events", controller.GetAllEvent)
		api.GET("/events/:event_id", controller.GetEventById)
		api.GET("/events/:community_id", controller.GetAllEventByCommunity)
		api.GET("/events/:event_category_id", controller.GetAllEventByCategory)

		// Route Event Participant Public
		api.GET("/event-participants", controller.GetAllEventParticipant)

		// Route Management Private
		secured := api.Group("/secured").Use(middleware.Auth())
		{
			// Route Users
			secured.GET("/users", controller.GetAllUser)
			secured.GET("/users/:user_id", controller.GetUserById)
			secured.POST("/users", controller.InsertUser)
			secured.PUT("/users/:user_id", controller.UpdateUser)
			secured.DELETE("/users/:user_id", controller.DeleteUser)

			// Route Communities
			secured.POST("/communities", controller.InsertCommunity)
			secured.PUT("/communities/:community_id", controller.UpdateCommunity)
			secured.DELETE("/communities/:community_id", controller.DeleteCommunity)

			// Route Event Categories
			secured.POST("/event-categories", controller.InsertEventCategory)
			secured.PUT("/event-categories/:event_category_id", controller.UpdateEventCategory)
			secured.DELETE("/event-categories/:event_category_id", controller.DeleteEventCategory)

			// Route Events
			secured.POST("/events", controller.InsertEvent)
			secured.PUT("/events/:event_id", controller.UpdateEvent)
			secured.DELETE("/events/:event_id", controller.DeleteEvent)

			// Route Event Participants
			secured.GET("/event-participants/:event_participant_id", controller.GeteventParticipantById)
			secured.POST("/event-participants", controller.InsertEventParticipant)
			secured.PUT("/event-participants/:event_participant_id", controller.UpdateEventParticipant)
			secured.DELETE("/event-participants/:event_participant_id", controller.DeleteEventParticipant)
		}
	}

	return router
}
