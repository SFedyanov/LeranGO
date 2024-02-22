package main

import (
	"fmt"
	"log"

	datafile "github.com/SFedyanov/grt/datafile"
)

func main() {
	lines, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lines)
}
