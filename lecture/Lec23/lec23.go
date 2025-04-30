/*
    1.Parameter vs Argument
	2.First Order Function
	  i.standard function
	  ii.Anonymous function
	  ii.Imediately Invoked function
	  iii.
	3.Higher Order Function
	  *Any one of the following 3 holds
	     i.function as a parameter
		 ii.return type as function
		 iii. both i,ii
*/

package main

import "fmt"

func processoperation(a int, b int, op func(x int, y int)) func(x int, y int) {
	op(a, b)

	return mul
}
func add(a int, b int) { //parameter => a,b
	c := a + b
	fmt.Println(c)
}
func mul(a int, b int) {
	c := a * b
	fmt.Println(c)
}
func call() func(x int, y int) {
	return add
}

func main() {
	m := processoperation(2, 5, add) //argument => 2,5

	m(5, 6)

	sum := call() //function expression

	sum(10, 20)
}
