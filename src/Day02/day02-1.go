/*
- 1 points for Rock (A) (X)
- 2 points for Paper (B) (Y)
- 3 points for Scissors (C) (Z)
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
	scanner := bufio.NewScanner(file)
	var strategy int
	// map scores/moves to points
	strategies := map[string]int{
		"B X": 1,
		"C Y": 2,
		"A Z": 3,
		"A X": 4,
		"B Y": 5,
		"C Z": 6,
		"C X": 7,
		"A Y": 8,
		"B Z": 9,
	}
	for scanner.Scan() {
		strategy += strategies[scanner.Text()]
	}
	fmt.Println(strategy)
}
