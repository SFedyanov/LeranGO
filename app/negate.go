package main

import "fmt"

func negate(myBoolean *bool) {
	fmt.Println(myBoolean)
	fmt.Println(*myBoolean)
	*myBoolean = !*myBoolean
}
func main() {
	truth := true
	negate(&truth)
	fmt.Println(truth)
	lies := false
	negate(&lies)
	fmt.Println(lies)
}
