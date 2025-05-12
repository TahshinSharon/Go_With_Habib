package main

import "fmt"

type User struct {
	Name   string
	Age    int
	Salary float64
}

// Pass by value
func printv(num [3]int) {
	fmt.Println(num)
}

//Pass by reference

func print(nums *[3]int) {
	fmt.Println(nums)
}

func prits(usr *User) { // Custom type as parameter
	fmt.Println(usr)
}

// Pointer or address of memory
func main() {

	x := 20

	p := &x // p is the address of x

	*p = 30 // Assign value into address of x

	fmt.Println("Value of x :", x)

	fmt.Println("Address of x :", p)

	fmt.Println("Value at the address of p :", *p)

	arr := [3]int{1, 2, 4}

	printv(arr)

	print(&arr)

	obj := User{
		Name:   "Tahshin",
		Age:    25,
		Salary: 30000.00,
	}

	prits(&obj)
}

/*
    Phases:
       1.Compilation Phase ()
	   2.Execution phase

go run lec25.go => compile it => lec25 => create binary file ./lec25 then run the binary file automatically
go build lec25.go => compile it => create ./lec25 file but don't run the binary file automatically


============= Compilation Phase ===========

***** Code Segement****
    print=func print(nums [3]int){..}
    main=func main(){.....}


***** Data Segment ****
   arr2
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
