package handlers

import (
	"net/http"

	"github.com/nick-mongkol/wfBackend/models"

	"github.com/gin-gonic/gin"
)

func GetAnnouncements(c *gin.Context) {
	// Fetch announcements (dummy data for now)
	announcements := []models.Announcement{
		{ID: 1, Title: "Warframe Alert", Description: "Special Alert Mission!"},
	}
	c.JSON(http.StatusOK, announcements)
}

func CreateAnnouncement(c *gin.Context) {
	var announcement models.Announcement
	if err := c.BindJSON(&announcement); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	// Here you would save the announcement to the database
	c.JSON(http.StatusCreated, announcement)
}
