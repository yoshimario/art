package main

import (
	"art/art-decoder/functions"
	"encoding/json"
	"flag"
	"html/template"
	"log"
	"net/http"
	"time"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

// Configuration flags
var (
	port            = flag.String("port", "8080", "Port to run the server on")
	enableAnimation = flag.Bool("animation", true, "Enable animations")
	debugMode       = flag.Bool("debug", false, "Enable debug mode")
	themeDefault    = flag.String("theme", "cyberpunk", "Default theme (cyberpunk, vaporwave, matrix)")
)

func main() {
	flag.Parse()

	// Configure routes
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/decoder", decoderHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Add the bonus API endpoint
	http.HandleFunc("/api/decode", apiDecodeHandler)
	http.HandleFunc("/api/encode", apiEncodeHandler)

	log.Printf("Server started on :%s", *port)
	log.Printf("Debug mode: %v", *debugMode)
	log.Printf("Default theme: %s", *themeDefault)
	
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Render the initial page with no result or error
		data := struct {
			Result      string
			Error       string
			Loading     bool
			Input       string
			Action      string
			Status      int
			Animation   bool
			DebugMode   bool
			ThemeDefault string
		}{
			Animation:   *enableAnimation,
			DebugMode:   *debugMode,
			ThemeDefault: *themeDefault,
		}
	
		templates.ExecuteTemplate(w, "index.html", data)
	}
}

func decoderHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        r.ParseForm()
        input := r.FormValue("input")
        action := r.FormValue("action")

        // Add artificial delay for visual feedback if debug mode is on
        if *debugMode {
            time.Sleep(1 * time.Second)
        }

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
            Result      string
            Error       string
            Loading     bool
            Input       string
            Action      string
            Status      int
            Animation   bool
            DebugMode   bool
            ThemeDefault string
        }{
            Result:      result,
            Loading:     false,
            Input:       input,
            Action:      action,
            Animation:   *enableAnimation,
            DebugMode:   *debugMode,
            ThemeDefault: *themeDefault,
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

// JSON response structure
type APIResponse struct {
    Success bool   `json:"success"`
    Result  string `json:"result,omitempty"`
    Error   string `json:"error,omitempty"`
}

// API endpoint for decoding
func apiDecodeHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    if r.Method != http.MethodPost {
        w.WriteHeader(http.StatusMethodNotAllowed)
        json.NewEncoder(w).Encode(APIResponse{
            Success: false,
            Error:   "Method not allowed, use POST",
        })
        return
    }

    // Parse JSON input
    var requestData struct {
        Input string `json:"input"`
    }

    err := json.NewDecoder(r.Body).Decode(&requestData)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(APIResponse{
            Success: false,
            Error:   "Invalid JSON input",
        })
        return
    }

    // Process the input
    result, err := functions.DecodeSingleLine(requestData.Input)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(APIResponse{
            Success: false,
            Error:   err.Error(),
        })
        return
    }

    // Return successful response
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(APIResponse{
        Success: true,
        Result:  result,
    })
}

// API endpoint for encoding
func apiEncodeHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    if r.Method != http.MethodPost {
        w.WriteHeader(http.StatusMethodNotAllowed)
        json.NewEncoder(w).Encode(APIResponse{
            Success: false,
            Error:   "Method not allowed, use POST",
        })
        return
    }

    // Parse JSON input
    var requestData struct {
        Input string `json:"input"`
    }

    err := json.NewDecoder(r.Body).Decode(&requestData)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(APIResponse{
            Success: false,
            Error:   "Invalid JSON input",
        })
        return
    }

    // Process the input
    result, err := functions.Encode(requestData.Input)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(APIResponse{
            Success: false,
            Error:   err.Error(),
        })
        return
    }

    // Return successful response
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(APIResponse{
        Success: true,
        Result:  result,
    })
}