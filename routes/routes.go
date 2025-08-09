// Package routes provides the HTTP routes for the chat application.
package routes

import (
	"chatapp/handlers"
	"chatapp/ws"

	"github.com/gin-gonic/gin"
)

func NewRouter(hub *ws.Hub) *gin.Engine {
	r := gin.Default()
	r.GET("/ws", handlers.ServeWs(hub))
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Chat server running")
	})
	return r
}
