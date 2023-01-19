package router

import (
	"FP-Sanbercode-Go-Batch-41/controller"
	"FP-Sanbercode-Go-Batch-41/middleware"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/", controller.GetCreator)

	api := router.Group("/api")
	{
		api.POST("/login", controller.GenerateToken)

		// Route Communities
		api.GET("/communities/:community_id", controller.GetCommunityById)
		api.GET("/communities", controller.GetAllCommunity)
		// api.POST("/communities", controller.InsertCommunity)
		// api.PUT("/communities/:community_id", controller.UpdateCommunity)
		// api.DELETE("/communities/:community_id", controller.DeleteCommunity)

		// Route Event Categories
		api.GET("/event-categories/:event_category_id", controller.GeteventCategoryById)
		api.GET("/event-categories", controller.GetAllEventCategory)
		// api.POST("/event-categories", controller.InsertEventCategory)
		// api.PUT("/event-categories/:event_category_id", controller.UpdateEventCategory)
		// api.DELETE("/event-categories/:event_category_id", controller.DeleteEventCategory)

		// Route Events
		api.GET("/events/:event_id", controller.GetEventById)
		api.GET("/events-by-category/:event_category_id", controller.GetAllEventByCategory)
		api.GET("/events-by-community/:community_id", controller.GetAllEventByCommunity)
		api.GET("/events", controller.GetAllEvent)
		// api.POST("/events", controller.InsertEvent)
		// api.PUT("/events/:event_id", controller.UpdateEvent)
		// api.DELETE("/events/:event_id", controller.DeleteEvent)

		// Route Event Participants
		api.GET("/event-participants/:event_participant_id", controller.GeteventParticipantById)
		api.GET("/event-participants", controller.GetAllEventParticipant)
		// api.POST("/event-participants", controller.InsertEventParticipant)
		// api.PUT("/event-participants/:event_participant_id", controller.UpdateEventParticipant)
		// api.DELETE("/event-participants/:event_participant_id", controller.DeleteEventParticipant)

		secured := api.Group("/secured").Use(middleware.Auth())
		{
			// Route Users
			secured.GET("/users/:user_id", controller.GetUserById)
			secured.GET("/users", controller.GetAllUser)
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
			secured.GET("/event-participants-by-event/:event_id", controller.GetAlleventParticipantByEvent)
			secured.GET("/event-participants", controller.GetAllEventParticipant)
			secured.POST("/event-participants", controller.InsertEventParticipant)
			secured.PUT("/event-participants/:event_participant_id", controller.UpdateEventParticipant)
			secured.DELETE("/event-participants/:event_participant_id", controller.DeleteEventParticipant)
		}
	}

	return router
}
