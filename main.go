package main

import (
	"fmt"
	"os"
	"strconv"

	"koodWordle/game"
	"koodWordle/io"
	"koodWordle/model"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <word index>")
		return
	}

	// Get the word index
	index, err := strconv.Atoi(os.Args[1])
	if err != nil || index < 0 {
		fmt.Println("Invalid word index")
		return
	}

	// Ask for username
	var username string
	fmt.Print("Enter your username: ")
	fmt.Scanln(&username)

	user := model.NewUser(username)

	// Load words from file
	words := io.LoadWords("wordle-words.txt")
	if len(words) == 0 {
		fmt.Println("No words available.")
		return
	}

	// Ensure index is valid
	if index >= len(words) {
		index = index % len(words)
	}

	secret := words[index]

	// Start the game
	game.Play(secret, user)

	// Ask for stats
	fmt.Print("Do you want to see your stats? (yes/no): ")
	var showStats string
	fmt.Scanln(&showStats)
	if showStats == "yes" {
		user.PrintStats()
	}
}
