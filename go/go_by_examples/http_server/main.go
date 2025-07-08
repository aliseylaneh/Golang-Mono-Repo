package main

import (
	"fmt"
	"io"
	"net/http"
)

func getHello(w http.ResponseWriter, request *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "Hello World")
}
func main() {
	http.HandleFunc("/", getHello)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println(err)
	}
}
