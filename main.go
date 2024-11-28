package main

import (
	AsciiArtWeb "AsciiArtWeb/ascii-art"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// `http.HandleFunc` handles GET and POST requests only via `/` and `/ascii-art`.
// server is started on `http://localhost:8080`.
func main() {
	http.HandleFunc("/", indexHandlerFunc)
	http.HandleFunc("/ascii-art", indexHandlerFunc)

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

// This function loads the HTML template (index.html) using tmpl.ParseFiles().
// This handles both GET and POST requests only.
// GET loads and renders the input form using the HTML template.
// POST reads user's input text (`textInputName`) and selected style (`bannerName`).
// Then calls `AsciiArt()` to convert the input text to ASCII art.
func indexHandlerFunc(w http.ResponseWriter, r *http.Request) {
	var tmpl *template.Template
	tmpl = template.Must(tmpl.ParseFiles("templates/index.html"))

	if r.URL.Path != "/" && r.URL.Path != "/ascii-art" {
		http.Error(w, "ERROR 404. Incorrect URL.", http.StatusNotFound)
		return
	}
	if r.Method == http.MethodPost {
		text := r.FormValue("textInputName")

		if text == "" { // if input is empty
			http.Error(w, "ERROR 400: Text input is empty.", http.StatusBadRequest)
			return
		}

		banner := r.FormValue("bannerName")
		AsciiArtWebOutput, err := AsciiArt(text, banner)
		if err != nil { // this handles all err we had on the terminal for ascii-art
			issue := "ERROR 500: " + err.Error()
			http.Error(w, issue, http.StatusInternalServerError)
			return
		}
		tmpl.ExecuteTemplate(w, "index.html", AsciiArtWebOutput)
	} else if r.Method == http.MethodGet {
		tmpl.ExecuteTemplate(w, "index.html", nil)
	} else { // if other method is used
		http.Error(w, "ERROR 405: Method Not Allowed.", http.StatusMethodNotAllowed)
		return
	}
}

// this handles the ascii-art generator
func AsciiArt(inputStr, bannerType string) (string, error) {

	var artSrcFilename string

	switch bannerType {
	case "standard":
		artSrcFilename = "standard.txt"
	case "shadow":
		artSrcFilename = "shadow.txt"
	case "thinkertoy":
		artSrcFilename = "thinkertoy.txt"
	}

	for _, c := range inputStr {
		if c != '\n' && c != '\r' {
			if c < ' ' || c > '~' {
				return "", fmt.Errorf("Character/s beyond standard ASCII printable characters code 32 to 126.")
			}
		}
	}

	return (AsciiArtWeb.ConvToArt(inputStr, artSrcFilename))

}
