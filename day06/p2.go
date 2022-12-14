package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	inputBytes, _ := ioutil.ReadFile("input.txt")
	input := string(inputBytes)

	charactersProcessed := 0
	for idx, _ := range input {
		inputMap := make(map[byte]bool)
		if idx+13 < len(input)-1 {
			for idx2 := idx; idx2 < idx+14; idx2++ {
				fmt.Println("here", string(input[idx2]))
				inputMap[input[idx2]] = true
			}
			fmt.Println(inputMap)
			if len(inputMap) == 14 {
				charactersProcessed += 14
				fmt.Println(charactersProcessed)
				return
			}
		}
		charactersProcessed++
	}

}
