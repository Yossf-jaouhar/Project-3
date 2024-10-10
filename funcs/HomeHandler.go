package funcs
import (
	"log"
	"net/http"
	"text/template"
)
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Only accept GET requests
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		http.ServeFile(w, r, "Templates/err400.html")
		return
	}
	// Ensure we only handle the root path "/"
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		http.ServeFile(w, r, "Templates/err404.html")
		return
	}
	// Parse the template
	tmpl, err := template.ParseFiles("Templates/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.ServeFile(w, r, "Templates/err500.html")
		log.Printf("Template parsing error: %v", err)
		return
	}
	// Send 200 OK status
	w.WriteHeader(http.StatusOK)
	data.ASCIIArt = ""
	// Render the template with the initialized (reset) data
	if err := tmpl.Execute(w, data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.ServeFile(w, r, "Templates/err500.html")
		log.Printf("Template execution error: %v", err)
	}
}
