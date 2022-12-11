package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input, _ := os.Open("input.txt")

	// close the file at end of program
	// defer is used to delay the execution of a function or a statement
	// until the nearby function returns
	defer input.Close()

	// Init buffer scanner I guess
	scanner := bufio.NewScanner(input)

	var curCalories, maxCalories, elfIdx, maxElfIdx int

	for scanner.Scan() {
		lineText := scanner.Text()
		if lineText == "" {
			if curCalories > maxCalories {
				maxCalories = curCalories
				maxElfIdx = elfIdx
			}
			elfIdx++
			curCalories = 0
			continue
		}
		cals, _ := strconv.Atoi(lineText)
		curCalories += cals

	}

	fmt.Println("Elf idx:", maxElfIdx, "Elf calories:", maxCalories)

}
