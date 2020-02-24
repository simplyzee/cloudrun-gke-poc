package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><head><title> Hello Cloud Run </title></head><body><p>Hello Cloud Run</p></body></html>\n")
}

func main() {
	port := os.Getenv("PORT")

	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+port, nil)
}
