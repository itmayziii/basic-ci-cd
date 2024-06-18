package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "Hello World - tommy")
	if err != nil {
		log.Printf("could not write response: %v", err)
	}
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("could not read body: %v\n", err)
	}

	_, err = w.Write(body)
	if err != nil {
		log.Printf("could not write response: %v\n", err)
	}
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/echo", echoHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		log.Printf("could not start server: %v\n", err)
		os.Exit(1)
	}
}
