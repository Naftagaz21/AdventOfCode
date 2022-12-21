package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Monkey struct {
	items []int
}

func main() {
	inputBytes, _ := ioutil.ReadFile("input.txt")
	inputLines := strings.Split(string(inputBytes), "\n\n")

	for i, line := range inputLines {
		fmt.Println(i, line)
		lines := strings.Split(line, "\n")
		startingItems := strings.Fields(lines[1])
		monkeyItems := make([]int, 0)

		// Get monkey items
		for _, item := range startingItems[2:] {
			x, _ := strconv.Atoi(strings.Trim(item, ","))
			monkeyItems = append(monkeyItems, x)
		}

		operationFields := strings.Fields(lines[2])
		//testField := strings.Fields(lines[2])
		//trueMonkey := strings.Fields(lines[3])
		//falseMonkey := strings.Fields(lines[4])
	}
}
