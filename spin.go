package main

import (
	"fmt"
	"math/rand"
)

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
