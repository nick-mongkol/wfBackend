package main

import (
	"fmt"
	"net/http"

	"github.com/nick-mongkol/wfBackend/config"
	"github.com/nick-mongkol/wfBackend/routes"

	"github.com/nick-mongkol/wfBackend/models"

	// Swagger docs

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan models.ChatMessage)

// WebSocket handler
func handleWebSocket(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Failed to upgrade connection:", err)
		return
	}
	defer ws.Close()

	clients[ws] = true
	for {
		var msg models.ChatMessage
		if err := ws.ReadJSON(&msg); err != nil {
			delete(clients, ws)
			break
		}
		broadcast <- msg
	}
}

func handleMessages() {
	for {
		msg := <-broadcast
		for client := range clients {
			if err := client.WriteJSON(msg); err != nil {
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func main() {
	// config.LoadEnv()

	r := gin.Default()

	// Swagger Documentation
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Setup routes
	routes.SetupRoutes(r)

	// WebSocket untuk chatroom
	r.GET("/ws", handleWebSocket)

	// Start broadcasting messages
	go handleMessages()

	// Start server
	port := config.GetEnv("PORT", "8080")
	r.Run(":" + port)
}
