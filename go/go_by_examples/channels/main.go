package main

import "fmt"

//func main() {
//	messages := make(chan string, 2)
//	go func() {
//		messages <- "Ping 2"
//		messages <- "Ping 3"
//	}()
//	for _ = range 2 {
//		fmt.Println(<-messages)
//	}
//}

func main() {
	messages := make(chan string, 2)

	messages <- "Ping 2"
	messages <- "Ping 3"

	for _ = range 2 {
		fmt.Println(<-messages)
	}
}
