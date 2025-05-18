package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"registerschemas/handlers"
	"syscall"
)

// enableCORS wraps an http.HandlerFunc and adds CORS headers to the response
func enableCORS(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		handler(w, r)
	}
}

func main() {
	// Create a channel to listen for signals
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	server := http.NewServeMux()

	// Wrap handlers with CORS middleware
	server.HandleFunc("/insert-schemas", enableCORS(handlers.InsertSchemas))
	server.HandleFunc("/fetch-schemas", enableCORS(handlers.FetchSchemas))
	server.HandleFunc("/delete-schema/{schemaId}", enableCORS(handlers.DeleteSchema))
	server.HandleFunc("/update-schema/{schemaId}", enableCORS(handlers.UpdateSchema))
	server.HandleFunc("/", enableCORS(handlers.HelloWorld))

	// Start the server in a goroutine
	go func() {
		log.Printf("Server starting on :8080")
		if err := http.ListenAndServe(":8080", server); err != nil {
			log.Fatal(err)
		}
	}()

	// Wait for termination signal
	<-sigs
	log.Println("Received shutdown signal, gracefully shutting down...")
}
