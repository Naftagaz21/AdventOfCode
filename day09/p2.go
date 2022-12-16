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

func calculateDistance(x2, y2, x1, y1 int) bool {
	x_dist := math.Pow(float64(x2-x1), 2)
	y_dist := math.Pow(float64(y2-y1), 2)
	distance := math.Round(math.Sqrt(x_dist + y_dist))
	if distance > 1 {
		return true
	} else {
		return false
	}

}

// 2597
func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()

	head := head{Coordinate{0, 0}}
	tail := tail{Coordinate{0, 0}, make([]Coordinate, 0)}

	scanner := bufio.NewScanner(input)

	headPositions := make([]Coordinate, 0)
	//headPositions = append(headPositions, head.location)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		direction := line[0]
		steps, _ := strconv.Atoi(line[1])

		tails := make([]Coordinate, 10)
		for i := 0; i < steps; i++ {
			oldHeadX := tails[0].x 
			oldHeadY := tails[0].y

			switch(direction) {
			case "U":
				tails[0].y += 1
			case "D":
				tails[0].y -= 1
			case "L":
				tails[0].x -= 1
			case "R":
				tails[0].x += 1

				for idx, coordVal := range(tails) {
					if idx == 0 {
						continue
					}
					if idx - 1 == 0 {
						
					}

					diff 
				}
			}
		}


		}
	}

	mp := make(map[Coordinate]bool, 0)
	for _, coord := range tail.visitedCoords {
		//fmt.Println(coord)

		mp[coord] = true
	}
	fmt.Println(len(mp) + 1)
	fmt.Println(tail.visitedCoords)

}
