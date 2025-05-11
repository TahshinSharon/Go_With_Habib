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
    Phases:
       1.Compilation Phase ()
	   2.Execution phase

go run lec25.go => compile it => lec25 => create binary file ./lec25 then run the binary file automatically
go build lec25.go => compile it => create ./lec25 file but don't run the binary file automatically


***** Code Segement****
   a=10
   main
   call25
   add

***** Data Segment ****
*/

/*
     *** Binary File ***

	 00110111101010101
	 10101010100101110
	 11100100010111011
	 10101010011100110
	 11001100011110001
	 11111011101010111
	 00110101010110111
	 10110011010111001
*/
