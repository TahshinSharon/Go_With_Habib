package main

import "fmt"

func add24(x int, y int) {
	z := x + y
	fmt.Println(z)
}

func main() {
	add24(5, 6)
	add24(12, 5)
}
func init() {
	fmt.Println("Hello")
}
