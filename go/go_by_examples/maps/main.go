package main

import (
	"fmt"
	"maps"
)

func main() {
	s := map[string]int{"Ali": 1}
	s["Ali"] = 5 // assign new values to a map key
	s["Seylaneh"] = 10
	fmt.Println(s)
	v := make(map[string]int) // create empty map using make
	fmt.Println(v)
	fmt.Println(v["Ali"]) // If you try to access a key which is not initialized in the map the zero value of the type will  be returned
	fmt.Println(len(v), len(s))
	delete(s, "Ali") // delete removes a key/value pair
	fmt.Println(s)
	//clear(s) // delete all the key/value pairs from a map
	fmt.Println(s)
	var valueExisted bool
	_, valueExisted = s["Seylaneh"]
	if valueExisted {
		fmt.Println("Value exists you can proceed.")
	}
	v = map[string]int{"Seylaneh": 10}
	if maps.Equal(s, v) {
		fmt.Printf("%v and %v are equal\n", s, v)
	}
}
