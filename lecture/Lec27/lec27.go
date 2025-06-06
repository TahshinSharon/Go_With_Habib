package main

import "fmt"

/*
    Phases:
       1.Compilation Phase ()
	   2.Execution phase

go run lec25.go => compile it => lec25 => create binary file ./lec25 then run the binary file automatically
go build lec25.go => compile it => create ./lec25 file but don't run the binary file automatically


============= Compilation Phase ===========

***** Code Segement****
   user

   main

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
type user struct {
	Name string
	Age  int
}

func main() {
	var user1 user //user1 is user  type variable
	user1 = user{
		Name: "Tahshin",
		Age:  26,
	}
	fmt.Println("User1 Name: ", user1.Name)
	fmt.Println("User1 Age: ", user1.Age)

	user2 := user{ // user type instance
		Name: "Sharon",
		Age:  36,
	}
	fmt.Println("User2 Name: ", user2.Name)
	fmt.Println("User2 Age: ", user2.Age)
}
