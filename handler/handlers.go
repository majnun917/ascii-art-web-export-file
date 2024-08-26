package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type FormData struct {
	Text     string
	Banner   string
	ASCIIArt string
}

var asciiArt string



// The main function that handles the HTTP methods (GET & POST)
func MainPage(w http.ResponseWriter, r *http.Request) {
	// Clear the global ASCII art variable
	asciiArt = ""

	switch r.Method {
	case "GET":
		tmpl, err := template.ParseFiles("./templates/index.html")
		if err != nil {
			http.Error(w, "Failed to parse file", http.StatusInternalServerError)
			return
		}
		data := FormData{}
		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	default:
		http.Error(w, "400: Method not allowed", http.StatusBadRequest)
		w.Write([]byte(http.StatusText(http.StatusBadRequest)))
	}
}

// Handler to process form submission
func SubmitForm(w http.ResponseWriter, r *http.Request) {
	const MaxLength = 5000

	switch r.Method {
	case "POST":
		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		text := r.FormValue("text")
		// Validate text length
		if len(text) > MaxLength {
			http.Error(w, fmt.Sprintf("Notice: Message exceeds maximum length of %d characters", MaxLength), http.StatusBadRequest)
			return
		}

		banner := r.FormValue("banner")
		if banner == "" {
			banner = "standard"
		}

		if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
			http.Error(w, "Notice: Please put a correct value: standard, shadow or thinkertoy", http.StatusBadRequest)
			return
		}

		art, err := PrintAsciiArt(text, banner) // Call the function to generate ASCII art
		if err != nil {
			if err.Error() == "Error" {
				http.Error(w, "Notice: only printable ASCII characters are allowed", http.StatusBadRequest)
				return
			}
		}

		asciiArt = art

		data := FormData{
			Text:     text,
			Banner:   banner,
			ASCIIArt: asciiArt,
		}

		tmpl, err := template.ParseFiles("./templates/ascii_art.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		tmpl.Execute(w, data)
		fmt.Println("Form successfully submitted !")
		fmt.Printf("Text: %s\nBanner: %s\n", text, banner)
	default:
		http.Error(w, "400: Method not allowed", http.StatusBadRequest)
		w.Write([]byte(http.StatusText(http.StatusBadRequest)))
	}
}

// Handler to download the ASCII art
func DownloadArt(w http.ResponseWriter, r *http.Request) {
	if asciiArt == "" {
		http.Error(w, "Notice: nothing to download, you didn't submit anything", http.StatusBadRequest)
		return
	}

	asciiArtBytes := []byte(asciiArt)

	w.Header().Set("Content-Disposition", "attachment; filename=ascii_art.txt")
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Length", strconv.Itoa(len(asciiArtBytes)))

	_, err := w.Write(asciiArtBytes)
	if err != nil {
		http.Error(w, "500: Failed to write file", http.StatusInternalServerError)
		return
	}
}
