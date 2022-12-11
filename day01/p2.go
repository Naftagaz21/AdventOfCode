package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()

	scanner := bufio.NewScanner(input)

	var maxCal1, maxCal2, maxCal3, curCals int

	for scanner.Scan() {
		line := scanner.Text()
		cals, err := strconv.Atoi(line)
		curCals += cals

		// Error occurs on newline apparently
		if err != nil {
			if curCals > maxCal3 {
				maxCal3 = curCals
			}
			if maxCal3 > maxCal2 {
				maxCal3, maxCal2 = maxCal2, maxCal3
			}
			if maxCal2 > maxCal1 {
				maxCal1, maxCal2 = maxCal2, maxCal1
			}

			curCals = 0
		}
	}

	fmt.Println(maxCal1 + maxCal2 + maxCal3)

}
