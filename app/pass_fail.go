// pass_fail
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var status string
	fmt.Print("Enter a grade:")
	grade, err := getFloat()
	if err != nil {
		log.Fatal(err)
	}
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}
	fmt.Println(status)
}

func getFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(input)
	return number, nil
}
