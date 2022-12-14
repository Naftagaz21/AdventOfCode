package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

func readStack(parts []string) []Stack {
	lines := strings.Split(parts[0], "\n")
	lines = lines[0 : len(lines)-1]
	for i, j := 0, len(lines)-1; i < j; i, j = i+1, j-1 {
		lines[i], lines[j] = lines[j], lines[i]
	}
	fmt.Println(lines)
	axe := lines[len(lines)-1]
	nr_of_slices := (len(axe)-1)/4 + 1

	twoD := make([]Stack, nr_of_slices)
	_ = twoD
	for _, line := range lines {
		//fmt.Println(idx, line)
		for b, i := 0, 1; i < len(line); i += 4 {
			char := string(line[i])
			if char != " " {
				twoD[b].Push(char)
			}
			b++
		}
	}

	return twoD
}

func moveStacks(itemPart string, stacks []Stack) []Stack {
	moveList := strings.Split(itemPart, "\n")
	blocksToMoveIdx, fromStackIdx, toStackIdx := 1, 3, 5
	var blocksToMove, fromStack, toStack int
	for _, move := range moveList {
		moves := strings.Fields(move)
		blocksToMove, _ = strconv.Atoi(moves[blocksToMoveIdx])
		fromStack, _ = strconv.Atoi(moves[fromStackIdx])
		toStack, _ = strconv.Atoi(moves[toStackIdx])

		fromStack--
		toStack--

		for i := 0; i < blocksToMove; i++ {
			item, _ := stacks[fromStack].Pop()
			stacks[toStack].Push(item)
			//fmt.Println(stacks)
		}
	}
	return stacks
}

func main() {
	text, _ := ioutil.ReadFile("input.txt")
	input := strings.TrimSuffix(string(text), "\n")
	parts := strings.Split(input, "\n\n")
	stacks := readStack(parts)
	//fmt.Println(parts[0])
	//fmt.Println(stacks)

	// TODO reverse parts[1]
	newStacks := moveStacks(parts[1], stacks)
	outStr := ""
	for _, x := range newStacks {
		el, _ := x.Pop()
		outStr += el
	}
	fmt.Println(newStacks)
	fmt.Println(outStr)
}
