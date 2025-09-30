package main

import (
	"fmt"
	"log"
)

func checkWin(spin [][]string, multipliers map[string]uint) []uint {
	lines := []uint{}

	for _, row := range spin {
		win := true
		checkSymbol := row[0]
		for _, symbol := range row[1:] {
			if symbol != checkSymbol {
				win = false
				break
			}
		}
		if win {
			lines = append(lines, multipliers[checkSymbol])
		} else {
			lines = append(lines, 0)
		}
	}
	return lines
}

func main() {
	symbols := map[string]uint{
		"A": 4,
		"B": 7,
		"C": 12,
		"D": 20,
	}

	multipliers := map[string]uint{
		"A": 20,
		"B": 10,
		"C": 5,
		"D": 2,
	}

	symbolSlice := generateSymbolArray(symbols)
	balance := uint(200)
	_, err := GetPlayerName()
	if err != nil {
		log.Fatal("An error has occurred: %w: ", err)
	}

	for balance > 0 {
		bet, err := GetBet(balance)
		if err != nil {
			log.Fatal("An error has occurred: %w", err)
		}
		if bet == 0 {
			break
		}
		balance -= bet
		spin := getSpin(symbolSlice, 3, 3)
		printSpin(spin)
		winningLines := checkWin(spin, multipliers)

		for i, multi := range winningLines {
			win := multi * bet
			balance += win
			if multi > 0 {
				fmt.Printf("Won $%d, (%dx) on Line #%d\n", win, multi, i+1)
			}
		}
	}
	fmt.Printf("You left with $%d\n", balance)
}
