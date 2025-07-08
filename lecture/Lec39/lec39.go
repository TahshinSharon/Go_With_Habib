package main

import "fmt"

func main() {
	var a int = 5  // int stores 64 bit init data in 64bit pc
	var b int8 = 7 // int8 stores 8 bit int data

	var x uint8 = 255 //uint8 stores unsigned 8 bit int data

	var j float64 = 10.23433 // if pc is 32 bit and variable is 64 bit then variable ocupied 2 32 bit cells from stack frame

	var flag bool = true // bool stores 8 bit
	fmt.Println(a, b, x, j, flag)
	r := '@'
	fmt.Println(r)
	fmt.Printf("%c", r)

	fmt.Printf("%d\n", a)
	fmt.Printf("%d\n", b)
	fmt.Printf("%d\n", x)
	fmt.Printf("%f\n", j)
	fmt.Printf("%v\n", flag)
	fmt.Printf("%T\n", flag)
}

/*
------------------------Vogus Data Types--------------------
   byte is alias for uint8-> which is 8 bit -> 1 byte

   rune is alias for int32 (unicode point)-> 32 bits -> 4bytes-> rune stores unicode charecter-> %c
*/
