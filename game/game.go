package game

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"koodWordle/model"
)

// ANSI colors
const (
	Green  = "\u001B[32m"
	Yellow = "\u001B[33m"
	Reset  = "\u001B[0m"
)

func Play(secret string, user *model.User) {
	fmt.Println("Welcome to Wordle! Guess the 5-letter word.")

	attempts := 6
	scanner := bufio.NewScanner(os.Stdin)

	for attempts > 0 {
		fmt.Printf("Enter your guess (%d attempts remaining): ", attempts)
		if !scanner.Scan() {
			break
		}
		guess := strings.ToLower(scanner.Text())
		if len(guess) != 5 {
			fmt.Println("Guess must be 5 letters.")
			continue
		}

		feedback := generateFeedback(secret, guess)
		fmt.Println("Feedback:", feedback)

		if guess == secret {
			fmt.Println(Green, "You win!", Reset)
			user.AddGame(secret, 6-attempts+1, true)
			return
		}

		attempts--
	}

	fmt.Println("You lost! Secret word was:", secret)
	user.AddGame(secret, 6, false)
}

// Generate feedback with ANSI colors
func generateFeedback(secret, guess string) string {
	result := ""
	for i := 0; i < len(secret); i++ {
		if guess[i] == secret[i] {
			result += Green + strings.ToUpper(string(guess[i])) + Reset
		} else if strings.Contains(secret, string(guess[i])) {
			result += Yellow + strings.ToUpper(string(guess[i])) + Reset
		} else {
			result += strings.ToUpper(string(guess[i]))
		}
	}
	return result
}
