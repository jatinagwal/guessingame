package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	secret := getRandomNumber()
	matching := false
	for !matching {

		guess := getUserInput()
		matching = isMatching(secret, guess)
	}
}

func isMatching(secret int, guess int) bool {
	if secret > guess {
		fmt.Println("Your guess is too small")
		return false
	} else if secret < guess {
		fmt.Println("Your guess is too big")
		return false
	} else {
		fmt.Println("YOU WON")
		return true
	}
}

func getUserInput() int {
	fmt.Println("Please type your guess: ")
	var input int
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Failed to get input")
	} else {
		fmt.Println("You guessed: ", input)
	}
	return input
}

func getRandomNumber() int {
	rand.Seed(time.Now().Unix())
	return rand.Int() % 11
}
