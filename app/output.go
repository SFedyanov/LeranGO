package main

import "fmt"

func main() {
	var pfloat float64 = 3.14151617181920
	var pint int = 456
	var pstr string = "str exampl"
	var pbool bool = true
	fmt.Printf("A float:\t %6.2f\t - %f\n", pfloat, pfloat)
	fmt.Printf("A fint: \t %6.2d\t - %d - %10d\n", pint, pint, pint)
	fmt.Printf("A fstr: \t %.2s\t - %20s - %s\n", pstr, pstr, pstr)
	fmt.Printf("A fbool:\t %t\t - %t\n", pbool, !pbool)
	fmt.Printf("A fany: \t %v\t - %v\t - %v\t - %v\n", pfloat, pint, pstr, pbool)
	fmt.Printf("A fany2: \t %#v\t - %#v\t - %#v\t - %#v\n", pfloat, pint, pstr, pbool)
	fmt.Printf("A fany3: \t %#v\t - %#v\t - %#v\t - %#v\n", "", "\t", "\n", "\\")
	fmt.Printf("A ftype: \t %T\t - %T\t - %T\t - %T\n", pfloat, pint, pstr, pbool)
	fmt.Printf("A fperc: \t %%\n")
	fmt.Println("")
	fmt.Printf("%12s | %s\n", "Column1", "Column2")
	fmt.Println("---------------------------")
	fmt.Printf("%12s | %d\n", "Stamps", 44)
	fmt.Printf("%12s | %d\n", "Column1", 55)
	fmt.Printf("%12s | %d\n", "Column1", 22)
}
