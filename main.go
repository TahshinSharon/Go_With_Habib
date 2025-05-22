package main

import "fmt"

// Generics
type stack[T any] struct {
	elements []T
}

func Print_With_Generics[T any, V string](items stack[T], name V) {
	fmt.Println(name, ":", items)
}

func main() {
	name := "Tahshin"

	myelemets := stack[int]{
		elements: []int{1, 2, 3, 4},
	}
	Print_With_Generics(myelemets, name)
}
