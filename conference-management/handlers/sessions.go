package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joshua468/conference-management/models"
	"github.com/joshua468/conference-management/storage"
)

func GetSessions(c *gin.Context) {
	sessions := []models.Session{}
	c.JSON(http.StatusOK, sessions)
}

func CreateSession(c *gin.Context) {
	var session models.Session
	if err := c.BindJSON(&session); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	storageSession := storage.Session{
		Title:       session.Title,
		Description: session.Description,
	}

	err := storage.CreateSession(&storageSession)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, storageSession)
}
