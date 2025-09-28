package main

import (
	"fmt"
	"log"
)

func getPlayerName() (string, error) {
	var name string
	fmt.Println("Welcome to Zel's Casino!")
	fmt.Println("Enter your name: ")
	_, err := fmt.Scanln(&name)
	if err != nil {
		return "", fmt.Errorf("ERR: %w", err)
	}
	fmt.Printf("Nice meeting you %s\n", name)
	return name, nil
}

func getBet(balance uint) (uint, error) {
	var bet uint
	for {
		fmt.Printf("Enter your bet (balance = $%d) or 0 to quit: ", balance)
		_, err := fmt.Scan(&bet)
		if err != nil {
			return 0, fmt.Errorf("ERR: %w", err)
		}
		if bet > balance {
			fmt.Println("You can not afford that, bet again.")
			continue
		}
		return bet, nil
	}

}

func main() {
	balance := uint(200)
	_, err := getPlayerName()
	if err != nil {
		log.Fatal("An error has occurred: %w: ", err)
	}

	for balance > 0 {
		bet, err := getBet(balance)
		if err != nil {
			log.Fatal("An error has occurred: %w", err)
		}
		if bet == 0 {
			break
		}
		balance -= bet
	}
	fmt.Printf("You left with $%d\n", balance)
}
