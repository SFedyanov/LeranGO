package main

import (
	"fmt"
	"log"

	datafile "github.com/SFedyanov/grt/datafile"
)

func main() {
	numbers, err := datafile.GetFloats("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	fmt.Println(sum / float64(len(numbers)))
}
