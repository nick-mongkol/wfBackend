package routes

import (
	"github.com/nick-mongkol/wfBackend/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/login", handlers.LoginHandler)
	r.GET("/announcements", handlers.GetAnnouncements)
	r.POST("/announcements", handlers.CreateAnnouncement)
}
