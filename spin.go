package main

import (
	"fmt"
	"math/rand"
)

// Function to generate a symbol array from a symbol-to-count map
func GenerateSymbolArray(symbols map[string]uint) []string {
	symbolArr := []string{}
	for symbol, count := range symbols {
		for i := uint(0); i < count; i++ {
			symbolArr = append(symbolArr, symbol)
		}
	}
	return symbolArr
}

// Function to generate a random number in a given range
func GetRandomNumber(min int, max int) int {
	return rand.Intn(max-min+1) + min
}

// Function to generate a spin result
func GetSpin(reel []string, rows int, cols int) [][]string {
	result := make([][]string, rows)
	for i := 0; i < rows; i++ {
		result[i] = []string{}
	}

	for col := 0; col < cols; col++ {
		selected := map[int]bool{}
		for row := 0; row < rows; row++ {
			for {
				randomIndex := GetRandomNumber(0, len(reel)-1)
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
func PrintSpin(spin [][]string) {
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
