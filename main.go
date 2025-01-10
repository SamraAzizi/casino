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
	for symbol, count := range symbols {
		for i := 0; i < int(count); i++ {

		}
	}
}

func main() {
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
