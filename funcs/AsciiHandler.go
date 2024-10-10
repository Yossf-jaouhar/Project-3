package funcs
import (
	"log"
	"net/http"
	"text/template"
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
		w.WriteHeader(http.StatusBadRequest)
		http.ServeFile(w, r, "Templates/err400.html")
		return
	}
	// Parse form data
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.ServeFile(w, r, "Templates/err500.html")
		log.Printf("Error parsing form data: %v", err)
		return
	}
	// Get and validate input text
	text := r.FormValue("inputText")
	if len(text) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		http.ServeFile(w, r, "Templates/err400.html")
		return
	}
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
	// Get the banner value, default to 'standard' if empty
	banner := r.FormValue("banner")
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
	tmpl, err := template.ParseFiles("Templates/index.html")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		http.ServeFile(w, r, "Templates/err404.html")
		log.Printf("Template parsing error: %v", err)
		return
	}
	if err := tmpl.Execute(w, data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.ServeFile(w, r, "Templates/err500.html")
		log.Printf("Template execution error: %v", err)
	}
}