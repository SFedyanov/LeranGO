package main

import "fmt"

var metersPerLiter = 10.0
var totalLiters float64

func main() {
	// var width, height, area float64
	// width = 4.2
	// height = 3.0
	// area = width * height
	// fmt.Printf("%.2f liters need\n", area/1.0)
	totalLiters += paintNeeded(4.2, 3.0)
	// width = 5.2
	// height = 3.5
	// area = width * height
	// fmt.Printf("%.2f liters need\n", area/1.0)
	totalLiters += paintNeeded(5.2, 3.5)
	totalLiters += paintNeeded(5.0, 3.3)
	fmt.Printf("%s %.0f (%.2f) \n", "Total liters:", totalLiters, totalLiters)
}

func paintNeeded(width float64, height float64) float64 {
	area := width * height
	liters := area / metersPerLiter
	fmt.Printf("%.2f liters need\n", liters)
	// totalLiters += liters
	return liters
}
