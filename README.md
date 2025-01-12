# Samra's Casino Game

## Overview
Samra's Casino Game is a simple command-line-based slot machine game written in Go. Players can place bets, spin the reels, and win based on the combinations of symbols. The game is interactive, allowing players to manage their balance and make strategic bets.

## Features
- Interactive gameplay with real-time input from players.
- Randomized spins for an authentic casino experience.
- Multiple symbol types with different payout multipliers.
- Balance management and betting system.

## Project Structure
The project is organized into three files for better modularity:

- **`main.go`**: Handles the core game loop, balance updates, and winning logic.
- **`spin.go`**: Contains the functionality for spinning the reels and generating results.
- **`utils.go`**: Includes utility functions for player interaction, such as getting names and bets.

## How to Run

### Prerequisites
- [Go](https://go.dev/) installed on your system (version 1.16 or higher).

### Steps
1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd <repository-folder>
   ```

2. Run the game:
   ```bash
   go run main.go spin.go utils.go
   ```

3. Follow the on-screen instructions to play the game.

## Gameplay Instructions
1. Start the game by entering your name.
2. You begin with a default balance of $200.
3. Enter your bet for each spin (bets cannot exceed your balance).
4. Spin the reels and view the results.
5. Win based on the multiplier of matching symbols on each line.
6. Continue playing until you quit or your balance reaches $0.

## Symbol and Multiplier Table
| Symbol | Count in Reel | Multiplier |
|--------|---------------|------------|
| A      | 4             | 20x        |
| B      | 7             | 10x        |
| C      | 12            | 7x         |
| D      | 20            | 2x         |

## Code Highlights
### Randomized Spins
The symbols on the reels are randomly selected using the `math/rand` package for a fair and unpredictable game experience.

### Winning Logic
The game checks for matching symbols on each line and calculates payouts based on predefined multipliers.

### Modular Design
The game logic is split across multiple files to improve readability and maintainability.

## Example Output
```
Welcome to Samra's Casino...
Enter your name: John
Welcome John, let's play!

Enter your bet, or 0 to quit (balance = $200): 50
A | B | C
D | D | D
A | A | A
Won $100, (20x) on Line #3
Balance: $250
```

