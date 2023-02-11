package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	//If we want to output something to the browser
	fmt.Fprintf(w, "<h1>Hello world</h1>")
}

func about(w http.ResponseWriter, r *http.Request) {
	//If we want to output something to the browser
	fmt.Fprintf(w, "<h1>About</h1>")
}

func main() {
	//It is like a router. It takes a route and you put a function to deal with that route.
	http.HandleFunc("/", index)

	http.HandleFunc("/about", about)

	fmt.Println("Server starting...")

	//in order for us to work or to serve:
	//El puerto y nil, que es como cualquier cosa
	http.ListenAndServe(":3000", nil)
}
