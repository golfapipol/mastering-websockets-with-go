package main

import (
	"context"
	"log"
	"net/http"

	"mastering-websockets-with-go/socket"
)

func main() {
	// Create a root ctx and a CancelFunc which can be used to cancel retentionMap goroutine
	rootCtx := context.Background()
	ctx, cancel := context.WithCancel(rootCtx)

	defer cancel()
	setupAPI(ctx)

	// Serve on port :8080, fudge yeah hardcoded port
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// setupAPI will start all Routes and their Handlers
func setupAPI(ctx context.Context) {
	// Create a Manager instance used to handle WebSocket Connections
	manager := socket.NewManager(ctx)

	// Serve the ./frontend directory at Route /
	http.Handle("/", http.FileServer(http.Dir("./frontend")))
	http.HandleFunc("/ws", manager.ServeWS)
}
