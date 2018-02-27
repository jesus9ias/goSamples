package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Profile struct for JSON response
type Profile struct {
	Name    string
	Hobbies []string
}

func handlerHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the web page: %s", r.URL.Path[1:])
}

func handlerJSON(w http.ResponseWriter, r *http.Request) {
	profile := Profile{"Pepe", []string{"snowboarding", "programming"}}

	jsResponse, err := json.Marshal(profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsResponse)
}

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/", handlerHome)
	server.HandleFunc("/json", handlerJSON)

	log.Fatal(http.ListenAndServe(":8080", server))
}
