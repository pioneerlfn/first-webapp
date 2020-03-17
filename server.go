package main

import (
	"fmt"
	"net/http"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.URL)
	fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path)
}
func main() {
	http.HandleFunc("/", handler)
	fmt.Println("server is running on 8080...")
	http.ListenAndServe(":8080", nil)
}
