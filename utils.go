package main

import "fmt"

func GetPlayerName() (string, error) {
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

func GetBet(balance uint) (uint, error) {
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
