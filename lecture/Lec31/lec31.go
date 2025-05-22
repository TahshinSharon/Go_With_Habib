package main

import "fmt"

func main() {

	arr := [4]string{"a", "ab", "abc", "abcd"}

	fmt.Println(arr)

	s := arr[1:3]

	fmt.Println(s)
}
