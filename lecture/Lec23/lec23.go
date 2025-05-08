/*
    1.Parameter vs Argument
	2.First Order Function
	  i.standard function
	  ii.Anonymous function
	  iii.Imediately Invoked function
	  iii. Function Expresiion
	3.Higher Order Function or first class function
	  *Any one of the following 3 holds
	     i.function as a parameter
		 ii.return type as function
		 iii. both i,ii
*/

package main

import "fmt"

func processOperation(a int, b int, op func(p int, q int)) { //Higher order Function as function takes function as parameter
	op(a, b)
}

func call() func(x int, y int) { //Higher order function as retuns a function
	return add
}

func add(x, y int) { // parameter
	c := x + y
	fmt.Println(c)
}
func main() {
	processOperation(2, 5, add) //The Function which we provide as argument or parameter in a function called Callback function

	sum := call() //function expression

	sum(4, 3)

	a := true //the data we assign into a variable called first class citizen so a function also called first class citizen in go

	fmt.Println(a)
}
