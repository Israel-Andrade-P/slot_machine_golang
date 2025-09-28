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

func main() {
	name, err := getPlayerName()
	if err != nil {
		log.Fatal("An error has occurred: %w", err)
	}
	fmt.Println(name)
}
