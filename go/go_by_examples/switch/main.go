package main

import (
	"fmt"
	"time"
)

func main() {
	i := 5

	switch i {
	case 2:
		fmt.Printf("The number is %v\n", i)
	case 3:
		fmt.Printf("The number is %v\n", i)
	default:
		fmt.Println("No value was recognized for this case.")
	}

	switch {
	case time.Now().Hour() < i:
		fmt.Printf("Time is less than %v.\n", i)
	default:
		fmt.Printf("Time is more than %v.\n", i)

	}

	switch time.Now().Weekday() {
	case time.Thursday, time.Friday:
		fmt.Printf("Today is off day work, %v\n", time.Now().Weekday())
	default:
		fmt.Printf("Today is work day, %v\n", time.Now().Weekday())
	}
	whatAmIFunction := func(value interface{}) {
		switch t := value.(type) {
		case bool:
			fmt.Println("I'm bool")
		case string:
			fmt.Println("I'm string")
		default:
			fmt.Printf("Don't know the type of %v\n", t)
		}
	}
	whatAmIFunction(true)
	whatAmIFunction(1)
	whatAmIFunction("Ali Seylaneh")

}
