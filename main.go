package main

import "fmt"

func getName() string {
	name := ""

	fmt.Println("Welcome To Samra's Casino...")
	fmt.Printf("Enter you name: ")
	_, err := fmt.Scanln(&name)
	if err != nil {
		return ""
	}
	fmt.Printf("Welcome %s, let's play!\n", name)
	return name
}

func getBet(balance uint) uint {

	var bet uint
	for true {
		fmt.Printf("Enter your bet, or 0 to quit (balance = $%d): ", balance)
		fmt.Scan(&bet)

		if bet > balance {
			fmt.Println("Bet cannot be larger than balance.")
		} else {
			break
		}

	}
	return bet
}

func generateSymbolArray(symbols map[string]uint) []string {
	symbolArr := []string{}
	for symbol, count := range symbols {
		for i := uint(0); i < count; i++ {
			symbolArr = append(symbolArr, symbol)

		}
	}
	return symbolArr
}

func getSpin(reel []string, row int, cols int) [][]string {
	result := [][]string{}

	for i := 0; i < rows; i++ {
		result = append(result, []string{})
	}

}

func main() {
	symbols := map[string]uint{
		"A": 4,
		"B": 7,
		"C": 12,
		"D": 20,
	}
	// multiplier := map[string]uint{
	// 	"A": 20,
	// 	"B" : 10,
	// 	"C" : 7,
	// 	"D" : 2,
	// }

	symbolArr := generateSymbolArray(symbols)
	fmt.Println(symbolArr)
	balance := uint(200)
	getName()

	for balance > 0 {
		bet := getBet(balance)
		if bet == 0 {
			break
		}
		balance -= bet
		fmt.Println(bet)

	}
	fmt.Printf("You left with, $%fd.\n", balance)

}
