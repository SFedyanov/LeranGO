// tocelsius преобразует температуру в градусах по Фаренгейту
// в градусы по Цельсию.

package main

import (
	"fmt"
	"log"

	keyboard "github.com/SFedyanov/keyboard"
)

func main() {
	fmt.Print("Enter a temperature in Fahrenheit: ")
	fahrenheit, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("%0.2f degrees Celsius\n", celsius)
}
