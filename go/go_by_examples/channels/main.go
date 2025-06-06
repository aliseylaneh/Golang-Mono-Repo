package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string)
	go channelReader(messages)

	messages <- "Ping"
	time.Sleep(time.Second)

}
func channelReader[T any](s chan T) {
	fmt.Println(<-s)
}
