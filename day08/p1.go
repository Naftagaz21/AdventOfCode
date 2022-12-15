package main

import (
	"bufio"
	"fmt"
	"os"
)

func traverseUp(treeHeight int, treeIdxX int, treeIdxY int, treeMap [][]int) bool {
	for yIdx := treeIdxY - 1; yIdx >= 0; yIdx-- {
		compareHeight := treeMap[yIdx][treeIdxX]
		if compareHeight >= treeHeight {
			return false
		}
	}
	return true
}

func traverseDown(treeHeight int, treeIdxX int, treeIdxY int, treeMap [][]int) bool {
	for yIdx := treeIdxY + 1; yIdx < len(treeMap); yIdx++ {
		compareHeight := treeMap[yIdx][treeIdxX]
		if compareHeight >= treeHeight {
			return false
		}
	}
	return true
}

func traverseLeft(treeHeight int, treeIdxX int, treeIdxY int, treeMap [][]int) bool {
	for xIdx := treeIdxX - 1; xIdx >= 0; xIdx-- {
		compareHeight := treeMap[treeIdxY][xIdx]
		if compareHeight >= treeHeight {
			return false
		}
	}
	return true
}

func traverseRight(treeHeight int, treeIdxX int, treeIdxY int, treeMap [][]int) bool {
	for xIdx := treeIdxX + 1; xIdx < len(treeMap[0]); xIdx++ {
		compareHeight := treeMap[treeIdxY][xIdx]
		if compareHeight >= treeHeight {
			return false
		}
	}
	return true
}

func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()

	scanner := bufio.NewScanner(input)

	oldSum := 0
	treeArr := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		oldSum += len(line)
		lineTrees := make([]int, 0)
		for _, char := range line {
			lineTrees = append(lineTrees, int(char-'0'))
		}
		treeArr = append(treeArr, lineTrees)
	}

	visibleTrees := 0
	iters := 0
	for treeY, treeLine := range treeArr {
		for treeX, tree := range treeLine {
			iters++
			check := traverseUp(tree, treeX, treeY, treeArr)
			if check {
				visibleTrees++
				continue
			}
			check = traverseDown(tree, treeX, treeY, treeArr)
			if check {
				visibleTrees++
				continue
			}
			check = traverseLeft(tree, treeX, treeY, treeArr)
			if check {
				visibleTrees++
				continue
			}
			check = traverseRight(tree, treeX, treeY, treeArr)
			if check {
				visibleTrees++
				continue
			}
		}
	}

	fmt.Println(visibleTrees, iters)
}
