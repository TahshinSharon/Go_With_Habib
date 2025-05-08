package main

import "fmt"

var sum = func(a, b int) { //Global Function Expression and assign function In variable
	s := a + b
	fmt.Println(s)
}

func main() {

	sum(4, 6)
	//add(1,2) this will not work
	// Function Expression Or Assign Function In Variable
	add := func(a int, b int) {
		c := a + b
		fmt.Println(c)
	}

	add(5, 7)

}
