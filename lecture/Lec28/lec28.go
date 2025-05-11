package main

import "fmt"

type User struct { //Create a Custom Type Named User
	Name string
	Age  int
}

var user1 = User{
	Name: "Tahshin",
	Age:  25,
}

func printUserDetails(usr User) { //Function takes Custom type user as parameter
	fmt.Println("Name: ", usr.Name)
	fmt.Println("Age: ", usr.Age)
}

func (usr User) printDetails() { //Receiver Function
	fmt.Println("Name: ", usr.Name)
	fmt.Println("Age: ", usr.Age)
}

func (usr User) call(a int) { // Receiver function with parameter
	fmt.Println("Name: ", usr.Name)
	fmt.Println("Age: ", a)
}

func main() {
	printUserDetails((user1)) //pass custom type User as Argument for any function

	user1.printDetails() //Call receiver function

	user1.call(10)
	user2 := User{
		Name: "Sharon",
		Age:  26,
	}
	printUserDetails(user2)
}

func init() {
	fmt.Println("*****Welcome To Lecture 28*****")
	fmt.Println()
}

/*
    Phases:
       1.Compilation Phase ()
	   2.Execution phase

go run lec25.go => compile it => lec25 => create binary file ./lec25 then run the binary file automatically
go build lec25.go => compile it => create ./lec25 file but don't run the binary file automatically


============= Compilation Phase ===========

***** Code Segement****

   User={
   Name string
   age int
   }
   init==func inti{...}

   main= func main(){...}

   printUserDetails=printUserDetails(usr User){...}

   printDetails=func(usr User)printDetails(){...}

   Call=func(usr User)call(a int){...}



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
