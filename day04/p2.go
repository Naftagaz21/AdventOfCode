package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func pairSplit(s string) []int {
	y := strings.Split(s, "-")
	values := make([]int, 0, len(y))
	for _, raw := range y {
		v, _ := strconv.Atoi(raw)
		values = append(values, v)
	}
	return values
}

func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()

	scanner := bufio.NewScanner(input)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		pair := strings.Split(line, ",")
		pairOne := pairSplit(pair[0])
		pairTwo := pairSplit(pair[1])

		if (pairOne[0] >= pairTwo[0] && pairOne[0] <= pairTwo[1]) ||
			(pairTwo[0] >= pairOne[0] && pairTwo[0] <= pairOne[1]) {
			sum++
		}
	}
	fmt.Println(sum)
}
