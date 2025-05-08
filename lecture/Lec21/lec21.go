package main

import "fmt"

// Expression: number of lines on code which is meaningfull for any task called expression
func main() { //main function
	fmt.Println("This is main function")

	func(a int, b int) { // Anonymous function: The function without any name
		c := a + b // summation and assigning expression
		fmt.Println(c)
	}(5, 7) //IIFE- Immediately Invoked Function Expression
}

func init() { //init function
	fmt.Println("This is init function")
}
