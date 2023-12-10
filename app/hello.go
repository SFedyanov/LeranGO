package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	hello()
	currentDate()
	replace()
}

func hello() {
	fmt.Println("Hello, Go!")
}

func currentDate() {
	var now time.Time = time.Now()
	var year int = now.Year()
	var month time.Month = now.Month()
	var day int = now.Day()
	fmt.Println(now)
	fmt.Println(year, month, day)
}

func replace() {
	brocken := "G# r#cks!"
	replacer := strings.NewReplacer("#", "o")
	fixed := replacer.Replace(brocken)
	fmt.Println(brocken, "\n", fixed)
}
