package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Coordinate struct {
	x int
	y int
}

type head struct {
	location Coordinate
}

type tail struct {
	location      Coordinate
	visitedCoords []Coordinate
}

func determineMove(x2, y2, x1, y1 int) (int, int) {
	//
	dx := math.Abs(float64(x2 - x1))
	dy := math.Abs(float64(y2 - y1))

	if dx <= 1 && dy <= 1 {
		return x1, y1
	} else if dx >= 2 && dy >= 2 {
		if y1 < y2 {
			y1 = y2 - 1
		} else {
			y1 = y2 + 1
		}

		if x1 < x2 {
			x1 = x2 - 1
		} else {
			x1 = x2 + 1
		}
	} else if dx >= 2 {
		if x1 < x2 {
			x1 = x2 - 1
		} else {
			x1 = x2 + 1
		}
		y1 = y2
	} else if dy >= 2 {
		x1 = x2
		if y1 < y2 {
			y1 = y2 - 1
		} else {
			y1 = y2 + 1
		}
	}
	return x1, y1

}

// 2597
func main() {
	input, _ := os.Open("input2.txt")
	defer input.Close()

	scanner := bufio.NewScanner(input)
	tailPositions := make([]Coordinate, 0)
	head := head{Coordinate{0, 0}}
	tails := make([]Coordinate, 9)

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		direction := line[0]
		steps, _ := strconv.Atoi(line[1])
		for i := 0; i < steps; i++ {
			switch direction {
			case "U":
				head.location.y += 1
			case "D":
				head.location.y -= 1
			case "L":
				head.location.x -= 1
			case "R":
				head.location.x += 1
			}
			tails[0].x, tails[0].y = determineMove(head.location.x, head.location.y, tails[0].x, tails[0].y)
			for k := 1; k < len(tails); k++ {
				tails[k].x, tails[k].y = determineMove(tails[k-1].x, tails[k-1].y, tails[k].x, tails[k].y)
				if k == 8 {
					tailPositions = append(tailPositions, Coordinate{tails[k].x, tails[k].y})
				}
			}
		}

	}

	mp := make(map[Coordinate]bool, 0)
	for _, coord := range tailPositions {
		mp[coord] = true
	}
	fmt.Println(len(tailPositions))
	fmt.Println(len(mp))

	//fmt.Println(tailPositions)

}
