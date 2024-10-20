package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func ViewSnippetHandler(w http.ResponseWriter, r *http.Request) {

	id, error := strconv.Atoi(r.URL.Query().Get("id"))

	if error != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Viewing the snippet id : %v", strconv.Itoa(id))
}

func CreateSnippetHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)

		return
	}

	w.Write([]byte("Creating snippet"))
}
