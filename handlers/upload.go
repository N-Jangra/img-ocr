package handlers

import (
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/otiai10/gosseract/v2"
)

// PageData represents data to be rendered in the index.html template
type PageData struct {
	ImageURL string // URL of uploaded image
	Text     string // Text extracted from the uploaded image via OCR
}

// Index handles HTTP GET requests to the root URL (/)
func Index(w http.ResponseWriter, r *http.Request) {
	// Render the index.html template with empty image and text
	renderTemplate(w, "", "")
}

// Upload handles HTTP POST requests for uploading images
func Upload(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST (should be)
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther) // Redirect to root URL if not POST
		return
	}

	// Get the uploaded file from the request
	file, header, err := r.FormFile("image")
	if err != nil { // Check for errors when reading the file
		http.Error(w, "Could not read image", http.StatusBadRequest)
		return
	}
	defer file.Close() // Ensure file is closed after use

	// Extract the file extension from the uploaded file's name
	ext := filepath.Ext(header.Filename)

	// Create a unique filename based on current time and file extension
	filename := time.Now().Format("20060102150405") + ext

	// Define the save path for the uploaded image (in the uploads directory)
	savePath := filepath.Join("uploads", filename)

	// Open the save path for writing the uploaded file contents
	out, err := os.Create(savePath)
	if err != nil { // Check for errors when creating the save path or writing to it
		http.Error(w, "Failed to save image", http.StatusInternalServerError)
		return
	}
	defer out.Close() // Ensure save path is closed after use

	// Copy the uploaded file contents to the save path
	io.Copy(out, file)

	// Initialize an OCR client and set its image source (the saved upload)
	client := gosseract.NewClient()
	defer client.Close()      // Close the client when done with it
	client.SetImage(savePath) // Set the image for OCR

	// Perform OCR on the uploaded image to extract text
	text, err := client.Text()
	if err != nil { // Check for errors during OCR (if any)
		text = "OCR failed" // Default text if OCR fails
	}

	// Render the index.html template with the extracted text and image URL
	renderTemplate(w, "/uploads/"+filename, text)
}

// renderTemplate renders an HTTP response from a given template file
func renderTemplate(w http.ResponseWriter, imageURL, text string) {
	// Parse the index.html template file (should be located in templates directory)
	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	// Create a PageData instance to pass to the template for rendering
	data := PageData{ImageURL: imageURL, Text: text}

	// Execute the rendered template on the HTTP response writer
	tmpl.Execute(w, data)
}
