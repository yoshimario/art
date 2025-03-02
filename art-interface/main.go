package main

import (
	"art/art-decoder/functions"
	"html/template"
	"log"
	"net/http"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/decoder", decoderHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Render the initial page with no result or error
		data := struct {
			Result  string
			Error   string
			Loading bool
			Input   string
			Action  string
			Status  int
		}{}
	
		templates.ExecuteTemplate(w, "index.html", data)
	}
}

func decoderHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        r.ParseForm()
        input := r.FormValue("input")
        action := r.FormValue("action")

        // Process the input
        var result string
        var err error
        var status int

        if action == "decode" {
            result, err = functions.DecodeSingleLine(input)
        } else if action == "encode" {
            result, err = functions.Encode(input)
        }

        // Prepare the data
        data := struct {
            Result  string
            Error   string
            Loading bool
            Input   string
            Action  string
            Status  int
        }{
            Result:  result,
            Loading: false,
            Input:   input,
            Action:  action,
        }

        if err != nil {
            // Set 400 status for malformed input
            status = http.StatusBadRequest
            data.Error = err.Error()
        } else {
            // Set 202 status for valid input
            status = http.StatusAccepted
        }
        
        data.Status = status
        w.WriteHeader(status)

        // Render the result
        templates.ExecuteTemplate(w, "index.html", data)
    } else {
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    }
}