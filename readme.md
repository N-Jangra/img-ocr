#
##  Go OCR Web App

A simple web application built with **Golang** that allows users to:

* Upload or drag-and-drop an image
* Extract text from the image using **OCR (Optical Character Recognition)**
* View the uploaded image and extracted text in the browser
* Copy the text manually (no JavaScript required)

---

##  What It Does

This application provides a web interface to upload image files (JPG, PNG, etc.). It uses the **Tesseract OCR engine** via the `gosseract` Go library to process the image and extract any readable text.

After submission:

1. The image is uploaded and stored in the `uploads/` directory.
2. OCR is performed server-side.
3. The webpage reloads and displays:

   * The uploaded image
   * Extracted text in a `<textarea>`

No client-side scripting (JavaScript) is required for image processing or rendering.

---

##  Features

*  Upload image through form
*  Displays uploaded image
*  Extracts text using Tesseract
*  Displays extracted text
*  Copy/paste manually

---

## 🛠️ Requirements

### 1. System Dependencies

You must have **Tesseract OCR** and its dependencies (Leptonica) installed:

####  Linux (Ubuntu/Debian)

```bash
sudo apt update
sudo apt install -y tesseract-ocr libtesseract-dev libleptonica-dev
```

####  macOS (with Homebrew)

```bash
brew install tesseract
```

####  Windows

* Download the installer from: [https://github.com/UB-Mannheim/tesseract/wiki](https://github.com/UB-Mannheim/tesseract/wiki)
* Add the installation path (e.g., `C:\Program Files\Tesseract-OCR`) to your system `PATH`
* Set `TESSDATA_PREFIX` to `C:\Program Files\Tesseract-OCR\tessdata`

---

### 2. Go Dependencies

```bash
go install
go get github.com/otiai10/gosseract/v2
```

---

## 🗂 Project Structure

```
go-ocr-app/
├── cmd/
│   └── server/
│       └── main.go              # Entry point
├── internal/
│   └── handlers/
│       └── upload.go            # Form and OCR logic
│   └── templates/
│       └── index.gohtml         # HTML rendering
├── static/
│   └── style.css                # CSS styles
├── uploads/                     # Uploaded image files
├── go.mod
└── go.sum
```

---

##  How to Run

### 1. Clone the Repo (or create the structure)

```bash
git clone https://github.com/N-Jangra/img-ocr
cd img-ocr
```

### 2. Initialize Go Module

```bash
go mod tidy
```

### 3. Run the Server

```bash
go run main.go
```

### 4. Open in Your Browser

Go to: [http://localhost:8080](http://localhost:8080)

---

##  How It Works

* The HTML form uploads the image via `POST /upload`
* Go handles the request:

  * Saves the image in `/uploads/`
  * Uses `gosseract` to extract text
  * Renders the same HTML template with image and text data
* The `uploads/` folder is statically served, so the image preview is just a `<img>` pointing to `/uploads/<filename>`

---

##  License

This project is open-source and free to use under the [MIT License](LICENSE).

#
