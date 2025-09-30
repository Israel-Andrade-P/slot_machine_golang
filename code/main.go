package main

import (
	"fmt"
	"log"
	"math/rand"
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

func generateSymbolArray(symbols map[string]uint) []string {
	symbolsSlice := make([]string, 0)

	for key, value := range symbols {
		for i := 0; i < int(value); i++ {
			symbolsSlice = append(symbolsSlice, key)
		}
	}
	return symbolsSlice
}

func getRandomNumber(min, max int) int {
	randomNum := rand.Intn(max-min+1) + min
	return randomNum
}

func getSpin(reel []string, rows, cols int) [][]string {
	result := make([][]string, 0)
	for i := 0; i < rows; i++ {
		result = append(result, []string{})
	}

	for col := 0; col < cols; col++ {
		selected := make(map[int]struct{})
		for row := 0; row < rows; row++ {
			for {
				randomIndex := getRandomNumber(0, len(reel)-1)
				if _, ok := selected[randomIndex]; !ok {
					result[row] = append(result[row], reel[randomIndex])
					selected[randomIndex] = struct{}{}
					break
				}
			}
		}
	}
	return result
}

func printSpin(spin [][]string) {
	for _, row := range spin {
		for j, symbol := range row {
			fmt.Printf(symbol)
			if j != len(row)-1 {
				fmt.Printf(" | ")
			}
		}
		fmt.Println()
	}
}

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
