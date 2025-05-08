package main

import "fmt"

var a = 10

func init() { //Init Function - You can not call this function,computer calls this automatically at first before main

	// this function can not take any parameter and don't return anything
	fmt.Println("I am the first function that execute first")

	fmt.Println(a)
	a = 20
	if a > 10 {
		a := 30 // variable shadowing
		fmt.Println(a)
	}
}
func main() {

	fmt.Println("Hello Main Function")

	fmt.Println(a)
}
