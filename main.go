package main

import (
	"fmt"
	"net/http"
)

func SayHello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
func main() {
	http.HandleFunc("/hello", SayHello)
	http.ListenAndServe(":8080", nil)
}
