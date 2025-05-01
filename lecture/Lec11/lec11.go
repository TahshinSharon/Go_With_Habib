package main

import "fmt"

// func add(n1 int, n2 int) int {
// 	sum := n1 + n2

//		return sum
//	}
func getNumbers(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	mul := n1 * n2
	return sum, mul
}
func main() {
	a := 10
	b := 20
	fmt.Println(getNumbers(a, b))
}
