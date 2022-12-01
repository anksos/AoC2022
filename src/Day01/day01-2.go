package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Read input from file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// Find top three maximum sums of calories
	maxThreeCalories := []int{0, 0, 0}
	currentElfCalories := 0

	for scanner.Scan() {
		// Get line from file
		line := scanner.Text()
		// Convert string to int
		calories, _ := strconv.Atoi(line)
		// Add calories to current elf
		currentElfCalories += calories
		// If is different from nil (NULL) then you have found an empty line
		if line == "" {
			// Check if currentElf has more calories than the maximum
			if currentElfCalories > maxThreeCalories[2] {
				maxThreeCalories[2] = currentElfCalories
			}
			if maxThreeCalories[2] > maxThreeCalories[1] {
				maxThreeCalories[2], maxThreeCalories[1] = maxThreeCalories[1], maxThreeCalories[2]
			}
			if maxThreeCalories[2] > maxThreeCalories[0] {
				maxThreeCalories[2], maxThreeCalories[0] = maxThreeCalories[0], maxThreeCalories[2]
			}
			// Start with new Elf
			currentElfCalories = 0
		}
	}
	// Print sum of maxThreeCalories
	fmt.Println(maxThreeCalories[0] + maxThreeCalories[1] + maxThreeCalories[2])
}
