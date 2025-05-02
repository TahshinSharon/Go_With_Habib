package main

import "fmt"

func main() {
	x := 20 // Local scope

	if x >= 18 {
		p := x //If block
		fmt.Println("You are ready to marry someone Less Than: ", p, " Years Old")
	}
}
