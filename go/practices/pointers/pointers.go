package main

import "fmt"

type Product struct {
	name    *string
	address string
}

func typeChecker(arg interface{}) {
	switch value := arg.(type) {
	case Product:
		fmt.Println("Type is product")
	case *Product:
		fmt.Println("Type is a pointer product")
	default:
		fmt.Println("Unknow type:", value)
	}
}
func failedUpdate(g *int) {
	x := 10
	g = &x
	fmt.Println(x)
	fmt.Println(g)
}
func update(px *int) {
	*px = 20

}

func makePointer[T any](t T) *T {
	return &t
}
func main() {
	product := &Product{name: makePointer("name"), address: "test"}
	fmt.Println(*product)
	typeChecker(product)
	//var g int = 10
	//failedUpdate(&g)
	//fmt.Println(g)
	//update(&g)
	//fmt.Println(g)
}
