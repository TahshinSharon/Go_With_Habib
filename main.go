package main

import (
	"fmt"
	"time"

	"github.com/thashinsharon/Go_with_Habib/lecture"
)

// Generics
type stack[T any] struct {
	elements []T
}

func Print_With_Generics[T any, V string](items stack[T], name V) {
	fmt.Println(name, ":", items)
}

func main() {
	name := "Tahshin"

	myelemets := stack[int]{
		elements: []int{1, 2, 3, 4},
	}
	Print_With_Generics(myelemets, name)

	boolCheck := lecture.Ptr(true)
	timeCheck := lecture.Ptr(time.Now().UTC())
	stringCheck := lecture.Ptr("Tahshin")

	fmt.Println("BoolPointer1:", boolCheck, "TimePointer1:", timeCheck, "stringPointer1:", stringCheck)
	boolCheck2 := lecture.BoolP(true)
	timeCheck2 := lecture.TimeP(time.Now().UTC())
	stringCheck2 := lecture.StringP("Tahshin")
	fmt.Println("BoolPointer2:", boolCheck2, "TimePointer2:", timeCheck2, "stringPointer2:", stringCheck2)

}
