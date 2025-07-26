package main

import "fmt"

/*
********* Clouse ************
********* Compile Phase ************
==== code segment: Stores compile time intems ====

	a=10
	outer
	outer-> anonymousfunc1()(show)

Data Segment: stores run time items

	stack autometic cleanup hoy

	* Only Show function can access money variable (they are binded in heap)

	* So when function end but we can still access some variable of this function for other anonymous function this is called clouser
*/
const a = 10

var p = 100

func outer() func() {
	money := 100 // escape analysis stores money into heap for future use

	age := 30

	fmt.Println("Age =", age)

	show := func() { //escape analysis also stores show function into heap memory
		money = money + a + p
		fmt.Println(money)
	}

	return show
}

func init() {
	fmt.Println("=== Clouser ===")
}

func call() {
	incr1 := outer() //incr1 stores show function

	incr1()
	incr1()

	incr2 := outer() //incr2 stores show function again newly escape analysis works again so money and show are again into heap
	incr2()
	incr2()
}

func main() {
	call()
}
