package main

import (
	"fmt"
	"time"
)

func printOut(name string) {
	for i := range 3 {
		fmt.Println(name, i)
	}
}
func main() {
	printOut("regular function")
	go printOut("goroutine")
	go func(msg string) {
		fmt.Println(msg)
	}("going")
	time.Sleep(time.Second)
	fmt.Println("Done")

}
