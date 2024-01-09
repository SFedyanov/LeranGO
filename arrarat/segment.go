package main

import "fmt"

func main() {
	// var notes []string
	notes := make([]string, 7)
	notes[0] = "do"
	notes[1] = "re"
	notes[6] = "mi"
	notes = append(notes, "do", "do")
	for index, values := range notes {
		fmt.Printf("%3v \t %3v\n", index, values)
	}
	primes := []int{
		2,
		3,
		4,
	}
	for index, values := range primes {
		fmt.Printf("%3v \t %3v\n", index, values)
	}
}
