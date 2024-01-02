package main

import "fmt"

var notes [7]string
var counters [10]int

func main() {
	notes[0] = "do"
	notes[1] = "re"
	notes[2] = "mi"
	notes[3] = "fa"
	fmt.Println("arrays:")
	// fmt.Println(notes[0])
	for i := 0; i < len(notes); i++ {
		// fmt.Println(notes[i])
		fmt.Printf("%4v: \"%v\" \t %T\n", i, notes[i], notes[i])
	}
	counters[0]++
	counters[0]++
	counters[2]++
	for i := 0; i < len(counters); i++ {
		// fmt.Println(notes[i])
		fmt.Printf("%4v: \"%v\" \t %T\n", i, counters[i], counters[i])
	}
	noteses := [7]string{
		"do",
		"re",
		"mi",
		"fa",
		"sol",
		"la",
		"si",
	}
	fmt.Println(noteses)
	fmt.Printf("%#v\n", noteses)
	for index, value := range noteses {
		fmt.Printf("%3v \t %v\n", index, value)
	}
}
