package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Delivery Hero Tech Hub Rocks!\n")
}

func main() {
	http.HandleFunc("/hello", hello)

	fmt.Println("Listening at 8090")
	http.ListenAndServe(":8090", nil)
}
