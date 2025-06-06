package main

import (
	"fmt"
)

// channelReader generic channel reader.
func channelReader[T any](s chan T) {
	fmt.Println(<-s)
}

// ping only accepts msg as our message
// and a channel as channel for sending values into it.
func ping(channel chan<- string, msg string) {
	channel <- msg
}

// pong shows us that how passed channels to functions,
// are supposed to work based on receiving or sending messages.
// when we say <-chan this means we can only accept values from the channel.
// when we say chan<- this means we can only send values to the channel.
func pong(channel <-chan string, pongs chan<- string) {
	msg := <-channel
	pongs <- msg
}
func main() {
	// This is an example of channel directions
	msg := "New message"
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, msg)
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
