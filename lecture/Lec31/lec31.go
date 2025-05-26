package main

import "fmt"

// variadic function
func print(numbers ...int) { ///slice as parameter
	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))
}

func main() {

	arr := [5]string{"a", "ab", "abc", "abcd", "abcde"} //Array

	fmt.Println(arr)

	s := arr[1:4] // Slice from array start from index 1 to just before index 4

	s1 := s[1:2] //slice from slice start from index 1 to just before index 2
	fmt.Println(s)
	fmt.Println(s1)

	sl := []int{1, 2, 3} // Slice Literal

	sl2 := make([]int, 3, 5) //Declare Slice using make function

	fmt.Println("slice:", sl, "len:", len(sl), "cap:", cap(sl))

	fmt.Println("slice:", sl2, "len:", len(sl2), "cap:", cap(sl2))

	var sl3 []int //empty slice or nil slice

	fmt.Println("slice:", sl3, "len:", len(sl3), "cap:", cap(sl3))
	sl3 = append(sl3, 1)
	sl3 = append(sl3, 1)
	fmt.Println("slice:", sl3, "len:", len(sl3), "cap:", cap(sl3))

	var x []int

	x = append(x, 1)
	x = append(x, 2)
	x = append(x, 3)

	y := x

	x[0] = 10
	x = append(x, 4)
	fmt.Println(x) //[0,2,3,4]
	y = append(y, 5)
	fmt.Println(x) //[10,2,3,5]
	fmt.Println(y) //[10,2,3,5]

	slice2()

	print(1, 2, 3, 4, 5) ///variadic function call using variable number of arguments

}

func changeSlice(v []int) []int {
	v[0] = 10
	v = append(v, 11)
	return v
}

func slice2() {
	x := []int{1, 2, 3, 4, 5}

	x = append(x, 6)
	x = append(x, 7)

	a := x[4:]

	y := changeSlice(a)

	fmt.Println(x)      //[1,2,3,4,10,6,7]
	fmt.Println(y)      //[10,6,7,11]
	fmt.Println(x[0:8]) //[1,2,3,4,10,6,7,11]
}

/*
   Slice Holds 3 Field Pointer, Length, Capacity

   Slice underlying array rule => 1024 -> 2 times/100% after that 25% inclease occure

*/
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


   ******* Stack *****
    stores executed function



	****** Heap *******


	stores escape analysus vaiables those can  be useful for future

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
