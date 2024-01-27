package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joshua468/conference-management/models"
	"github.com/joshua468/conference-management/storage"
)

func GetAttendees(c *gin.Context) {
	attendees := []models.Attendee{}
	c.JSON(http.StatusOK, attendees)
}

func CreateAttendee(c *gin.Context) {
	var attendee models.Attendee
	if err := c.BindJSON(&attendee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convert models.Attendee to storage.Attendee
	storageAttendee := storage.Attendee{
		Name: attendee.Name,
	}

	err := storage.CreateAttendee(&storageAttendee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, storageAttendee)
}
