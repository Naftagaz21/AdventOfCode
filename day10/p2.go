package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func checkCycle(valueArr []int, cycle int, register int) []int {
	signalStrenght := cycle * register
	if cycle == 20 {
		valueArr[0] = signalStrenght
	} else if cycle == 60 {
		valueArr[1] = signalStrenght
	} else if cycle == 100 {
		valueArr[2] = signalStrenght
	} else if cycle == 140 {
		valueArr[3] = signalStrenght
	} else if cycle == 180 {
		valueArr[4] = signalStrenght
	} else if cycle == 220 {
		valueArr[5] = signalStrenght
	}
	return valueArr
}

func main() {
	inputBytes, _ := ioutil.ReadFile("input.txt")
	inputLines := strings.Split(string(inputBytes), "\n")

	lineIdx, register, addRegister := 0, 1, 0
	cycleStop := false
	valueArr := make([]int, 6)
	secondCycle := 0
	for cycle := 2; cycle <= 1000; cycle++ {
		if (cycle-2)%40 == 0 {
			fmt.Println("")
			secondCycle = 0
		}

		if secondCycle == register-1 || secondCycle == register || secondCycle == register+1 {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
		secondCycle++

		if lineIdx > len(inputLines)-1 {
			break
		}
		lineFields := strings.Fields(inputLines[lineIdx])
		if len(lineFields) == 0 {
			break
		}
		if cycleStop {
			cycleStop = false
			register += addRegister
			lineIdx++
			valueArr = checkCycle(valueArr, cycle, register)
			continue
		} else if lineFields[0] == "noop" {
			lineIdx++
			valueArr = checkCycle(valueArr, cycle, register)
			continue
		} else if lineFields[0] == "addx" {
			cycleStop = true
			x, _ := strconv.Atoi(lineFields[1])
			addRegister = x
			valueArr = checkCycle(valueArr, cycle, register)
			continue
		}
	}
	sum := 0
	for _, val := range valueArr {
		sum += val
	}
	fmt.Println(sum)
}
