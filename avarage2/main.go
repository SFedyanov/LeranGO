package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func average(numbers ...float64) float64 {
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	return sum / float64(len(numbers))
}

func main() {
	arguments := os.Args[1:]
	// var sum float64 = 0
	var numbers []float64
	for _, value := range arguments {
		number, err := strconv.ParseFloat(value, 64)
		if err != nil {
			log.Fatal(err)
		}
		// sum += number
		numbers = append(numbers, number)
	}
	// sampleCount := float64(len(arguments))
	// fmt.Printf("Average: %0.2f\n", sum/sampleCount)
	fmt.Printf("Average: %0.2f\n", average(numbers...))
}
