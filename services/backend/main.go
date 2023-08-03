package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

var port = flag.String("port", "8080", "Port number to run the server on")

func main() {

	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World! %s", time.Now())
	})
	addr := ":" + *port

	log.Printf("Server is running on port %s\n", *port)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Error starting the server: %v\n", err)
	}
}
