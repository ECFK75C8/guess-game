package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX = 3

func main() {
	fmt.Println("Guess the number!")

	// generate a random number
	source := rand.NewSource(time.Now().UnixNano())
	randomizer := rand.New(source)
	secretNumber := randomizer.Intn(10)

	var guess int

	for i := 0; i < MAX; i++ {
		fmt.Printf("Please input your guess (%d times left!)", (MAX - i))
		fmt.Scan(&guess)

		if guess > secretNumber {
			fmt.Println("Too big")
		} else if guess < secretNumber {
			fmt.Println("Too small")
		} else {
			fmt.Println("Correct guess!")
			break
		}

		if i == 2 {
			fmt.Println("Game ended!")
		}
	}
}
