package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Read input file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// Find maximum sum of calories
	maxCalories := 0
	currentElfCalories := 0
	for scanner.Scan() {
		// Get line from file
		line := scanner.Text()
		// Convert string to int
		calories, _ := strconv.Atoi(line)
		// Add calories to current elf
		currentElfCalories += calories
		// Check if current elf has more calories than the maximum
		if currentElfCalories > maxCalories {
			maxCalories = currentElfCalories
		}
		// check if line is empty
		if line == "" {
			// Start with new Elf
			currentElfCalories = 0
		}
	}
	// Print maxCalories
	fmt.Println(maxCalories)
}
