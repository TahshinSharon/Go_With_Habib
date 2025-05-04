package main

import (
	"fmt"

	"github.com/thashinsharon/Go_with_Habib/lecture/Lec16/mathlib"
)

/*
// Today key point

	*go mod init
	* Function of other package must be start with cpital letter otherwise it could not found
*/
func main() {
	fmt.Println("Check Summation:")
	mathlib.Add(10, 20)

	fmt.Println("Check Multiplication:")

	mathlib.Mul(10, 20)
}
