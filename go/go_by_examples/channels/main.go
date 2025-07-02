package main

import (
	"fmt"
	"time"
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

// Chanel synchronization example.
// worker This function is under processing.
// take a channel to notify another goroutine that this function is done.
func worker(done chan<- bool) {
	fmt.Print("working... ")
	time.Sleep(2 * time.Second)
	fmt.Println("Done.")
	done <- true
}
func selectPractice() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "one"
	}()
	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "two"
	}()
	for range 2 {
		select {
		case msg1 := <-ch1:
			fmt.Println("received", msg1)
		case msg2 := <-ch2:
			fmt.Println("received", msg2)
		}
	}
}
func main() {
	// This is an example of channel directions
	//msg := "New message"
	//pings := make(chan string, 1)
	//pongs := make(chan string, 1)
	//ping(pings, msg)
	//pong(pings, pongs)
	//fmt.Println(<-pongs)
	//done := make(chan bool, 1)
	//go worker(done)
	//<-done
	selectPractice()
}
