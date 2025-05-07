package main

import "fmt"

// variable shadowing

var a = 10 // here a is 10

func main() {
	age := 30

	if age > 18 {
		a := 47 //now a is 47 so this is variable shadowing
		fmt.Println(a)
	}

	fmt.Println(a) //here a is 10
}
