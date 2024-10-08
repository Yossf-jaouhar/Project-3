package funcs

import (
	"log"
	"net/http"
)

type PageData struct {
	ASCIIArt string
	Text     string
	Banner   string
}

var data = &PageData{Banner: "standard"}

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	// Only accept POST requests
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		http.ServeFile(w, r, "Templates/err405.html")
		return
	}

	// Parse the form data and handle potential errors
	if err := r.ParseForm(); err != nil {

		w.WriteHeader(http.StatusBadRequest)
		http.ServeFile(w, r, "Templates/err400.html")
		// Log the detailed error message for debugging
		log.Printf("Error parsing form data: %v", err)
		return
	}

	// Get and validate input text
	text := r.FormValue("inputText")
	if len(text) > 2000 {
		w.WriteHeader(http.StatusBadRequest)
		http.ServeFile(w, r, "Templates/err400.html")
		return
	}
	for _, char := range text {
		if !(char >= 32 && char <= 126 || char == '\n' || char == '\r') {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "Templates/err400.html")
			return
		}
	}
	if text == "" {
		w.WriteHeader(http.StatusBadRequest)
		http.ServeFile(w, r, "Templates/err400.html")
		return
	}

	// Get the banner value, default to 'standard' if empty
	banner := r.FormValue("banner")
	if banner == "" {
		banner = "standard"
	}

	// Read the banner data
	bannerData, err := ReadB(banner)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.ServeFile(w, r, "Templates/err500.html")
		log.Printf("Error reading banner: %v", err)
		return
	}

	// Generate ASCII art from the input text
	asciiArt := TreatData(bannerData, text)
	data.ASCIIArt = asciiArt
	data.Text = ""
	data.Banner = banner

	http.Redirect(w, r, "/", http.StatusFound)
}
