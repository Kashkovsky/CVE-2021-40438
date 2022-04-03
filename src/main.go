package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Request from: %q", html.EscapeString(r.RemoteAddr))
		fmt.Println("Request from: ", html.EscapeString(r.RemoteAddr))
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}
