package main

import "fmt"

/* Standard or named Function
* Anonymous function
* Function Expression or Assign Function in Variable
* Higher Order Functin or First Class Function
* Callback Function
* Variadic function
* Init Function- You can not call this , Computer calls thsi automatically
* Closure - close over
* Defer Function -last in firt out
* Receiver Function
* IIFE- Immediately invoked function expression
 */

func add(a int, b int) { //Standard or Named Function
	sum := a + b
	fmt.Println(sum)
}
func main() {
	add(4, 7)
}
