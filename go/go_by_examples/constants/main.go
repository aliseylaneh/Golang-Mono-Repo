package main

import (
	"fmt"
	"math/rand"
)

const s string = "constant"
const ApiFlagSuccess string = "success"

func main() {
	var ApiFlag string
	randomNumber := rand.Int()
	if randomNumber > 10 {
		ApiFlag = ApiFlagSuccess
	}
	a := "new constant"
	var b = a
	if ApiFlag == "success" {
		fmt.Println(b)
	}
}
