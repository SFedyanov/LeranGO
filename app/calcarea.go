package main

import (
	"fmt"
)

var metersPerLiter = 10.0
var totalLiters float64
var err error
var amount float64

func main() {
	amount, err = paintNeeded(4.2, 3.0)
	if err != nil {
		fmt.Println(err)
	} else {
		totalLiters += amount
	}

	amount, err = paintNeeded(5.2, 3.5)
	if err != nil {
		fmt.Println(err)
	} else {
		totalLiters += amount
	}

	amount, err = paintNeeded(5.0, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		totalLiters += amount
	}

	amount, err = paintNeeded(5.0, -3.3)
	if err != nil {
		fmt.Println(err)
	} else {
		totalLiters += amount
	}

	fmt.Printf("%s %.0f (%.2f) \n", "Total liters:", totalLiters, totalLiters)
}

func paintNeeded(width float64, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("a width of %0.2f is invalid", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("a height of %0.2f is invalid", height)
	}
	area := width * height
	liters := area / metersPerLiter
	fmt.Printf("%.2f liters need\n", liters)
	// totalLiters += liters
	return liters, nil
}
