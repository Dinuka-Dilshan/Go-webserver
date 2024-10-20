package main

import (
	"log"
	"net/http"
)

func main() {

	serveMux := http.NewServeMux()

	serveMux.HandleFunc("/", HomeHandler)
	serveMux.HandleFunc("/snippet/view", ViewSnippetHandler)
	serveMux.HandleFunc("/snippet/create", CreateSnippetHandler)

	error := http.ListenAndServe(":5000", serveMux)

	if error != nil {
		log.Fatal("Error starting server: ", error)
	}
}
