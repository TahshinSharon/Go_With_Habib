package main

import "fmt"

func main() {
	age := 20
	if age > 18 {
		fmt.Println("You are eligible to be marrid")
	} else if age < 18 {
		fmt.Println("You are not eligible to be marrid but you can love someone")
	} else {
		fmt.Println("You are just a teenager,not eligible to be married")
	}
	sex := "Male"

	if age > 18 && sex == "Male" {
		fmt.Println("You Have to Marry a girl")
	}

	day := 1

	switch day {
	case 7, 1:
		fmt.Println("Off Day")

	default:
		fmt.Println("Working Day")

	}
}
