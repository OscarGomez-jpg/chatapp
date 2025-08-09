package main

import (
	"chatapp/client"
	"chatapp/server"
	"flag"
	"fmt"
	"os"
)

func main() {
	mode := flag.String("mode", "", "server or client")
	addr := flag.String("addr", "localhost:8080", "address")
	flag.Parse()

	switch *mode {
	case "server":
		srv := server.NewServer()
		fmt.Println("Server listening on", *addr)
		if err := srv.Listen(*addr); err != nil {
			fmt.Println("Server error:", err)
		}
	case "client":
		if err := client.Start(*addr); err != nil {
			fmt.Println("Client error:", err)
		}
	default:
		fmt.Println("Usage: go run main.go -mode=server|client [-addr=localhost:8080]")
		os.Exit(1)
	}
}
