// pass_fail
package main

import (
	"fmt"
	"log"

	keyboard "github.com/SFedyanov/keyboard"
)

func main() {
	var status string
	fmt.Print("Enter a grade:")
	grade, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}
	fmt.Println("A grade of", grade, "is", status)
}
