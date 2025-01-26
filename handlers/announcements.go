package handlers

import (
	"net/http"

	"github.com/nick-mongkol/wfBackend/models"

	"github.com/gin-gonic/gin"
)

func GetAnnouncements(c *gin.Context) {
	// Fetch announcements (dummy data for now)
	// announcements := []models.Announcement{
	// 	{ID: 1, Title: "Warframe Alert", Description: "Special Alert Mission!"},
	// }
	c.JSON(http.StatusOK, announcements)
}

var announcements = []models.Announcement{
	{ID: 1, Title: "New Event", Description: "Join us for a special event!"},
	{ID: 2, Title: "Maintenance", Description: "Server maintenance on Saturday."},
}

func CreateAnnouncement(c *gin.Context) {
	var announcement models.Announcement
	if err := c.BindJSON(&announcement); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	announcements = append(announcements, announcement)
	// Here you would save the announcement to the database
	c.JSON(http.StatusCreated, announcements)
}
