/*
- 1 points for Rock (A)
- 2 points for Paper (B)
- 3 points for Scissors (C)
- X means need to lose
- Y means need to draw
- Z means need to win
- 0 points if player lost
- 3 points if it's a draw
- 6 points if it's a win
- Winner is the player with the most points from the whole tournament
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Read input file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file")
	}
	defer file.Close()
	var strategy int
	// map scores/moves to points and outcome
	strategies := map[string]int{
		"B X": 1, // lose
		"C X": 2, // lose
		"A X": 3, // lose
		"A Y": 4, // draw
		"B Y": 5, // draw
		"C Y": 6, // draw
		"C Z": 7, // win
		"A Z": 8, // win
		"B Z": 9, // win
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		strategy += strategies[scanner.Text()]
	}
	fmt.Println(strategy)
}
