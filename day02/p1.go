package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()

	scanner := bufio.NewScanner(input)

	// A, X - Rock
	// B, Y - Paper
	// C, Z - Scissors

	// TODO research map
	scores := map[string]int{"A X": 4, "A Y": 8, "A Z": 3, "B X": 1, "B Y": 5, "B Z": 9, "C X": 7, "C Y": 2, "C Z": 6}
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		sum += scores[line]
	}
	fmt.Println(sum)
}
