package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const s = "สวัสดี"
	fmt.Println(s, len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x", s[i])
	}
	fmt.Println()
	fmt.Println(utf8.RuneCountInString(s))
	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}
	// Loop over string
	someString := "Ali Seylaneh"
	for _, i2 := range someString {
		fmt.Printf("%v\n", i2)
	}
	fmt.Println(utf8.RuneCountInString(someString))
}
