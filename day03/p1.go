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

	for scanner.Scan() {
		line := scanner.Text()
		line_len := len(line)
		line_break := line_len / 2
		first_part := line[0:line_break]
		second_part := line[line_break:line_len]

		for _, char := range first_part {
			trigger := false
			for _, char2 := range second_part {
				//fmt.Println(char, char2)
				if char == char2 {
					fmt.Println(string(char), first_part, second_part, valMap[int(char)])
					sum += valMap[int(char)]
					trigger = true
					break
				}
			}

			if trigger == true {
				break
			}
		}

	}
	fmt.Println(sum)
}
