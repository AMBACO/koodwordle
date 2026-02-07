package model

import "fmt"

type User struct {
	Username string
	Games    int
	Wins     int
	Attempts int
}

func NewUser(username string) *User {
	return &User{
		Username: username,
	}
}

func (u *User) AddGame(word string, attempts int, win bool) {
	u.Games++
	u.Attempts += attempts
	if win {
		u.Wins++
	}
	// Here you can also write to stats.csv if needed
}

func (u *User) PrintStats() {
	avg := 0.0
	if u.Games > 0 {
		avg = float64(u.Attempts) / float64(u.Games)
	}
	fmt.Println("Stats for", u.Username)
	fmt.Println("Games played:", u.Games)
	fmt.Println("Games won:", u.Wins)
	fmt.Printf("Average attempts per game: %.2f\n", avg)
}
