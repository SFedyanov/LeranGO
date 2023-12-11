package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileInfo, err := os.Stat("./app/hello.go")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fileInfo.Size())
	fmt.Println(fileInfo.Mode())
}
