package main

import (
	"fmt"
	"net/http"
)

func index_handler(a http.ResponseWriter, c *http.Request) {
	fmt.Fprintf(a, "GO is a cool language")
}

func call_handler(a http.ResponseWriter, c *http.Request) {
	fmt.Fprintf(a, "here is contact us page")
}

func about_handler(a http.ResponseWriter, c *http.Request) {
	fmt.Fprintf(a, "here is about us page using golang")
}

func main() {
	http.HandleFunc("/home", index_handler)
	http.HandleFunc("/contact-us", call_handler)
	http.HandleFunc("/about-us", about_handler)
	http.ListenAndServe(":8000", nil)
}
