package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joshua468/conference-management/models"
	"github.com/joshua468/conference-management/storage"
)

func GetConferences(c *gin.Context) {
	conferences := []models.Conference{}
	c.JSON(http.StatusOK, conferences)
}

func CreateConference(c *gin.Context) {
	var conference models.Conference
	if err := c.BindJSON(&conference); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	storageConference := storage.Conference{
		Name: conference.Name,
	}

	err := storage.CreateConference(&storageConference)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, storageConference)
}
