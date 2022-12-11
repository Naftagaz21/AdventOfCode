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

	// A Rock
	// B Paper
	// C Scissors

	// X - Lose
	// Y - Draw
	// Z - Win

	// TODO research map
	shapeOut := map[string]string{
		"A X": "A C",
		"A Y": "A A",
		"A Z": "A B",
		"B X": "B A",
		"B Y": "B B",
		"B Z": "B C",
		"C X": "C B",
		"C Y": "C C",
		"C Z": "C A"}
	scores := map[string]int{"A A": 4, "A B": 8, "A C": 3, "B A": 1, "B B": 5, "B C": 9, "C A": 7, "C B": 2, "C C": 6}

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		line = shapeOut[line]
		sum += scores[line]
		fmt.Println(line, scores[line])
	}
	fmt.Println(sum)
}
