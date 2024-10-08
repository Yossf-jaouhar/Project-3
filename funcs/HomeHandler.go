package funcs

import (
	"log"
	"net/http"
	"text/template"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
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

	// Initialize empty data to pass to the template

	w.WriteHeader(http.StatusOK)

	// Render the template with the initialized data
	if err := tmpl.Execute(w, data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.ServeFile(w, r, "Templates/err500.html")

		log.Printf("Template execution error: %v", err)
	}
}
