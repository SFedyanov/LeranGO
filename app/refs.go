package main

import (
	"fmt"
)

func main() {
	// amount := 6
	// amount2 := 6
	// fmt.Println(amount)
	// fmt.Println(&amount)
	// fmt.Println(amount2)
	// fmt.Println(&amount2)
	// var myInt int
	// fmt.Println(reflect.TypeOf(&myInt))
	// var myFloat float64
	// fmt.Println(reflect.TypeOf(&myFloat))
	// var myBool bool
	// fmt.Println(reflect.TypeOf(&myBool))

	// myInt := 4
	// myIntPointer := &myInt
	// fmt.Println(myInt)
	// fmt.Println(myIntPointer)
	// fmt.Println(*myIntPointer)
	// *myIntPointer = 5
	// fmt.Println(myInt)
	// fmt.Println(myIntPointer)
	// fmt.Println(*myIntPointer)
	// myInt = 6
	// fmt.Println(myInt)
	// fmt.Println(myIntPointer)
	// fmt.Println(*myIntPointer)
	// myIntPointer2 := &myIntPointer
	// fmt.Println(myIntPointer)
	// fmt.Println(myIntPointer2)
	// fmt.Println(*myIntPointer2)

	var myInt int
	var myIntPointer *int
	myInt = 42
	myIntPointer = &myInt
	fmt.Println(*myIntPointer)
}
