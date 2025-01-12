package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Function to generate a symbol array from a symbol-to-count map
func generateSymbolArray(symbols map[string]uint) []string {
	symbolArr := []string{}
	for symbol, count := range symbols {
		for i := uint(0); i < count; i++ {
			symbolArr = append(symbolArr, symbol)
		}
	}
	return symbolArr
}

// Function to generate a random number in a given range
func getRandomNumber(min int, max int) int {
	return rand.Intn(max-min+1) + min
}

// Function to generate a spin result
func getSpin(reel []string, rows int, cols int) [][]string {
	result := make([][]string, rows)
	for i := 0; i < rows; i++ {
		result[i] = []string{}
	}

	for col := 0; col < cols; col++ {
		selected := map[int]bool{}
		for row := 0; row < rows; row++ {
			for {
				randomIndex := getRandomNumber(0, len(reel)-1)
				if !selected[randomIndex] {
					selected[randomIndex] = true
					result[row] = append(result[row], reel[randomIndex])
					break
				}
			}
		}
	}
	return result
}

// Function to print the spin results
func printSpin(spin [][]string) {
	for _, row := range spin {
		for j, symbol := range row {
			fmt.Print(symbol)
			if j != len(row)-1 {
				fmt.Print(" | ")
			}
		}
		fmt.Println()
	}
}

// Function to check for winning lines
func checkWin(spin [][]string, multipliers map[string]uint) []uint {
	lines := []uint{}

	for _, row := range spin {
		win := true
		checkSymbol := row[0]
		for _, symbol := range row[1:] {
			if checkSymbol != symbol {
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

// Main function
func main() {
	rand.Seed(time.Now().UnixNano())

	symbols := map[string]uint{
		"A": 4,
		"B": 7,
		"C": 12,
		"D": 20,
	}
	multipliers := map[string]uint{
		"A": 20,
		"B": 10,
		"C": 7,
		"D": 2,
	}

	symbolArr := generateSymbolArray(symbols)
	balance := uint(200)
	GetName()

	for balance > 0 {
		bet := GetBet(balance)
		if bet == 0 {
			break
		}

		balance -= bet
		spin := getSpin(symbolArr, 3, 3)
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
	fmt.Printf("You left with $%d.\n", balance)
}
