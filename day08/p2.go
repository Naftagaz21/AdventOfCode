package main

import (
	"bufio"
	"fmt"
	"os"
)

func traverseUp(treeHeight int, treeIdxX int, treeIdxY int, treeMap [][]int) int {
	treeCount := 0
	for yIdx := treeIdxY - 1; yIdx >= 0; yIdx-- {
		compareHeight := treeMap[yIdx][treeIdxX]
		if compareHeight >= treeHeight {
			treeCount++
			return treeCount
		}
		treeCount += 1
	}
	return treeCount
}

func traverseDown(treeHeight int, treeIdxX int, treeIdxY int, treeMap [][]int) int {
	treeCount := 0
	for yIdx := treeIdxY + 1; yIdx < len(treeMap); yIdx++ {
		compareHeight := treeMap[yIdx][treeIdxX]
		if compareHeight >= treeHeight {

			treeCount++
			return treeCount
		}
		treeCount++
	}
	return treeCount
}

func traverseLeft(treeHeight int, treeIdxX int, treeIdxY int, treeMap [][]int) int {
	treeCount := 0
	for xIdx := treeIdxX - 1; xIdx >= 0; xIdx-- {
		compareHeight := treeMap[treeIdxY][xIdx]
		if compareHeight >= treeHeight {
			treeCount++
			return treeCount
		}
		treeCount++
	}
	return treeCount
}

func traverseRight(treeHeight int, treeIdxX int, treeIdxY int, treeMap [][]int) int {
	treeCount := 0
	for xIdx := treeIdxX + 1; xIdx < len(treeMap[0]); xIdx++ {
		compareHeight := treeMap[treeIdxY][xIdx]
		if compareHeight >= treeHeight {
			treeCount++
			return treeCount
		}
		treeCount++
	}
	return treeCount
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

	//visibleTrees := 0
	iters := 0
	maxTreeCount := 0
	for treeY, treeLine := range treeArr {
		var count1, count2, count3, count4 int
		for treeX, tree := range treeLine {
			iters++
			count1 = traverseUp(tree, treeX, treeY, treeArr)
			count2 = traverseDown(tree, treeX, treeY, treeArr)
			count3 = traverseLeft(tree, treeX, treeY, treeArr)
			count4 = traverseRight(tree, treeX, treeY, treeArr)

			// TODO learn error handling in GO
			total := 0
			if count1 == 0 || count2 == 0 || count3 == 0 || count4 == 0 {
				total = 1
			} else {
				total = count1 * count2 * count3 * count4
			}
			if total > maxTreeCount {
				fmt.Println(count1, count2, count3, count4)
				maxTreeCount = total
			}
		}

	}

	fmt.Println(maxTreeCount)
}
