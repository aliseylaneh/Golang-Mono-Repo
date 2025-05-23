package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base
	str string
}
type describer interface {
	describe() string
}

func main() {
	co := container{
		base: base{
			12,
		},
		str: "Ali Container",
	}
	fmt.Printf("Base number is %v\n", co.describe())
	fmt.Printf("Access directly to the embbeded variable %v\n", co.num)
	fmt.Printf("Access directly from the embedded variable %v\n", co.base.num)

	var d describer = co
	fmt.Println("describer:", d.describe())

}
