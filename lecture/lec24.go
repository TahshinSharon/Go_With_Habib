package lecture

import "fmt"

func add24(x int, y int) {
	z := x + y
	fmt.Println(z)
}

func Lec24() {
	add24(5, 6)
	add24(12, 5)
}
func init() {
	fmt.Println("Hello")
}
