package lecture

import "fmt"

type user struct {
	Name string
	Age  int
}

func Lec27() {
	var user1 user
	user1 = user{
		Name: "Tahshin",
		Age:  26,
	}
	fmt.Println("User1 Name: ", user1.Name)
	fmt.Println("User1 Age: ", user1.Age)

	user2 := user{
		Name: "Sharon",
		Age:  26,
	}
	fmt.Println("User2 Name: ", user2.Name)
	fmt.Println("User2 Age: ", user2.Age)
}
