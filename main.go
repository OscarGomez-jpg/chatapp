package main

import (
	"chatapp/handlers"
	"chatapp/ws"

	"github.com/gin-gonic/gin"
)

func main() {
	hub := ws.NewHub()
	go hub.Run()

	r := gin.Default()
	r.GET("/ws", handlers.ServeWs(hub))
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Chat server running")
	})

	r.Run(":8080")
}
