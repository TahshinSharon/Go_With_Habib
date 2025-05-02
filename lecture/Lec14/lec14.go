package main

import "fmt"

var (
	a = 20 //global scope
	b = 30
)

func add(x int, y int) {
	z := x + y //local scope
	fmt.Println(z)
}

func main() {
	p := 30 //local scope
	q := 40
	add(p, q)
	add(a, b)
	add(a, p)
}
