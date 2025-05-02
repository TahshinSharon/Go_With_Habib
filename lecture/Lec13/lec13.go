package main

import "fmt"

func printwelcomemessage() {
	fmt.Println("Welcome to the application")
}
func getUserName() string {
	var name string
	fmt.Println("Enter Your Name")
	fmt.Scanln(&name)
	return name
}
func add(n1 int, n2 int) {
	sum := n1 + n2
	fmt.Println("Summation Is:", sum)
}
func printstring(str string) {
	fmt.Println("Welcome: ", str)
}
func getTwoNumbers() (int, int) {
	var n1, n2 int
	fmt.Println("Enter First Number")
	fmt.Scanln(&n1)
	fmt.Println("Enter Secound Number")
	fmt.Scanln(&n2)
	return n1, n2
}
func main() {
	printwelcomemessage()
	name := getUserName()
	printstring(name)
	n1, n2 := getTwoNumbers()
	add(n1, n2)
	fmt.Println("Bye Bye: ", name)
}
