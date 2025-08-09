// Package handlers for chat server
package handlers

import (
	"chatapp/ws"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func ServeWs(hub *ws.Hub) gin.HandlerFunc {
	return func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			return
		}
		hub.Register <- conn

		go func() {
			defer func() {
				hub.Unregister <- conn
			}()
			for {
				_, msg, err := conn.ReadMessage()
				if err != nil {
					break
				}
				hub.Broadcast <- msg
			}
		}()
	}
}
