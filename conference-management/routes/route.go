package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/joshua468/conference-management/handlers"
)

func SetupRoutes(router *gin.Engine) {
	attendeeGroup := router.Group("/attendees")
	{
		attendeeGroup.GET("/", handlers.GetAttendees)
		attendeeGroup.POST("/", handlers.CreateAttendee)
	}

	conferenceGroup := router.Group("/conferences")
	{
		conferenceGroup.GET("/", handlers.GetConferences)
		conferenceGroup.POST("/", handlers.CreateConference)
	}

	sessionGroup := router.Group("/sessions")
	{
		sessionGroup.GET("/", handlers.GetSessions)
		sessionGroup.POST("/", handlers.CreateSession)
	}
}
