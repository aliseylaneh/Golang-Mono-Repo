package main

import (
	"fmt"
)

func main() {
	messages := make(chan string)
	go func() {
		messages <- "Ping 2"
	}()
	fmt.Println(<-messages)
}
