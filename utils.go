


// Function to get the user's name
func GetName() string {
	var name string

	fmt.Println("Welcome to Samra's Casino...")
	fmt.Printf("Enter your name: ")
	_, err := fmt.Scanln(&name)
	if err != nil {
		fmt.Println("Error reading input. Please try again.")
		return ""
	}
	fmt.Printf("Welcome %s, let's play!\n", name)
	return name
}

// Function to get the user's bet
func GetBet(balance uint) uint {
	var bet uint
	for {
		fmt.Printf("Enter your bet, or 0 to quit (balance = $%d): ", balance)
		_, err := fmt.Scan(&bet)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			continue
		}
		if bet > balance {
			fmt.Println("Bet cannot be larger than balance.")
		} else {
			break
		}
	}
	return bet
}
