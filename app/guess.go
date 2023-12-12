// guess - игра, в которой нужно угадать случайное число от 1 до 100
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(100) + 1
	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")
	fmt.Println("You have 10 attempts.")
	reader := bufio.NewReader(os.Stdin)
	success := false
	for guesses := 1; guesses <= 10; guesses++ {
		fmt.Println("Attempt", guesses)
		fmt.Print("Make a guess: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(guess)
		if guess < target {
			fmt.Println("Opps. Your guess was LOW.")
		} else if guess > target {
			fmt.Println("Oops. Your guess was HIGH.")
		} else if guess == target {
			success = true
			fmt.Println("Good job! You guessed it! It is:", target)
			break
		}
	}
	if !success {
		fmt.Println("Sorry, you didn't guess my number. It was:", target)
	}
	fmt.Println("Thank you!")
	fmt.Println("bye!")
	fmt.Print("Prees Enter to Exit")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
}
