package main

import "fmt"

/*
*Code Segment: Stores all function and const variable (Stores Read Only code which are not changeable)

*Data Segment: Stores all global variable(variables which can changeable)

* Stack: When a function start executing a memory alocates from stack for this specific function called stack frame

Every segment dynamic based on size of the code
*/
var a = 10

func add(x int, y int) {
	z := x + y
	fmt.Println(z)
}

func main() {
	add(5, 6)
	add(a, 5)
}
func init() {
	fmt.Println("Hello")
}
