package main

import (
	"chatapp/routes"
	"chatapp/ws"
)

func main() {
	hub := ws.NewHub()
	go hub.Run()

	r := routes.NewRouter(hub)
	r.Run(":8080")
}
