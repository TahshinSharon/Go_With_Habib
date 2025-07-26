package main

import "fmt"

func calculate() (result int) {
	fmt.Println("First", result)
	show := func() {
		result = result + 10
		fmt.Println("Defer", result)
	}

	defer show()
	result = 5
	fmt.Println("Second", result)
	return
}
func cal() int {
	result := 0
	fmt.Println("First", result)
	show := func() {
		result = result + 10
		fmt.Println("Defer", result)
	}

	defer show()
	result = 5
	fmt.Println("Second", result)
	return result
}
func magicbox() (result int) {
	fmt.Println("First", result)
	show := func() {
		result = result + 10
		fmt.Println("Defer", result)
	}

	defer show()
	result = 5

	p := func(a int) {
		fmt.Println("ami", a)
	}
	defer p(result)
	defer fmt.Println(result)
	fmt.Println("Second", result)
	defer fmt.Println(5)
	return
}
func main() {
	fmt.Println("Main First", calculate())
	b := cal()
	fmt.Println("Main Second", b)

	c := magicbox()

	fmt.Println("Main Third", c)
}
