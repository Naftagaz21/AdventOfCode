package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()

	valMap := make(map[int]int)

	for charIdx, charVal := 97, 1; charIdx <= 122; charIdx++ {
		//fmt.Println(charIdx, charVal)
		valMap[charIdx] = charVal
		charVal++

	}

	for charIdx, charVal := 65, 27; charIdx <= 90; charIdx++ {
		//fmt.Println(charIdx, charVal)
		valMap[charIdx] = charVal
		charVal++
	}

	scanner := bufio.NewScanner(input)

	sum := 0

	var groupLines []string
	for scanner.Scan() {
		line := scanner.Text()
		groupLines = append(groupLines, line)
		if len(groupLines) == 3 {
			for _, char := range groupLines[0] {
				trigger := false
				for _, char2 := range groupLines[1] {
					for _, char3 := range groupLines[2] {
						if char == char2 && char2 == char3 {
							sum += valMap[int(char)]
							trigger = true
							break
						}
					}
					if trigger {
						break
					}
				}
				if trigger {
					break
				}
			}
			groupLines = nil
		}

	}
	fmt.Println(sum)
}
