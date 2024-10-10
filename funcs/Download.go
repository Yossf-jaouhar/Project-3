package funcs

import (
	"net/http"
	"strconv"
)

// Download handles file download requests.
func Download(w http.ResponseWriter, r *http.Request) {
	// Check if the method is POST.
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		http.ServeFile(w, r, "Templates/err405.html")
		return
	}

	// Parse form data.
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.ServeFile(w, r, "Templates/err400.html")
		return
	}

	// Get the form value.
	result := r.FormValue("arttext")
	if result == "" {
		w.WriteHeader(http.StatusBadRequest)
		http.ServeFile(w, r, "Templates/err400.html")
		return
	}

	// Set headers to trigger file download.
	w.Header().Set("Content-Disposition", "attachment; filename=result.txt")
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Length", strconv.Itoa(len(result)))

	// Write the content to the response, triggering the download.
	w.Write([]byte(result))
}
