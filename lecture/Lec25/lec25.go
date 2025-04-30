package main

import "fmt"

const a = 10

var p = 100

func call25() {
	add := func(x int, y int) {
		z := x + y
		println(z)
	}

	add(5, 6)
	add(p, a)
}

func main() {
	call25()
	fmt.Println(a)
}

/*
    1.Compilation Phase ()
	2.Execution phase

go run main.go => compile it => main => create binary file ./main then run the binary file automatically
go build main.go => compile it => create ./main file but don't run the binary file automatically
*/
