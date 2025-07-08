package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func getHello(w http.ResponseWriter, request *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "Hello World")
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", getHello)
	err := http.ListenAndServe(":8000", mux)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("Server closed.")
	} else if err != nil {
		fmt.Printf("error starting server:%s\n", err)
		os.Exit(1)
	}
}
