package main

import "fmt"

var arr2 = [3]string{"I", "Love", "You"}

func main() {

	arr := [2]int{3, 6}
	fmt.Println(arr)

	fmt.Println(arr2[1])

}
func init() {
	fmt.Println("*****Welcome To Lecture 29*****")
}

/*
    Phases:
       1.Compilation Phase ()
	   2.Execution phase

go run lec25.go => compile it => lec25 => create binary file ./lec25 then run the binary file automatically
go build lec25.go => compile it => create ./lec25 file but don't run the binary file automatically


============= Compilation Phase ===========

***** Code Segement****
   init==func inti{...}

   main= func main(){...}



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
