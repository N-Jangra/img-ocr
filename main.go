package main

import (
	"fmt"
	"net/http"
	"ocr/handlers"
)

func main() {
	// Static file handler
	// Serve static files from 'static' directory
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Upload file handler
	// Serve uploaded files from 'uploads' directory
	http.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("uploads"))))

	// Route handlers
	// Handle root URL ('/') with Index handler function
	http.HandleFunc("/", handlers.Index)

	// Handle '/upload' URL with Upload handler function
	http.HandleFunc("/upload", handlers.Upload)

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
