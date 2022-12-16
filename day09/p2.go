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
//2597
func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()

	head := head{Coordinate{0, 0}}
	tail1 := tail{Coordinate{0, 0}, make([]Coordinate, 0)}
	tail2 := tail{Coordinate{0, 0}, make([]Coordinate, 0)}
	tail3 := tail{Coordinate{0, 0}, make([]Coordinate, 0)}

	tail4 := tail{Coordinate{0, 0}, make([]Coordinate, 0)}
	tail5 := tail{Coordinate{0, 0}, make([]Coordinate, 0)}
	tail6 := tail{Coordinate{0, 0}, make([]Coordinate, 0)}

	tail7 := tail{Coordinate{0, 0}, make([]Coordinate, 0)}
	tail8 := tail{Coordinate{0, 0}, make([]Coordinate, 0)}
	tail9 := tail{Coordinate{0, 0}, make([]Coordinate, 0)}

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		direction := line[0]
		steps, _ := strconv.Atoi(line[1])

		for i := 0; i < steps; i++ {
			oldHeadY := head.location.y
			oldHeadX := head.location.x
			switch direction {
			case "U":
				head.location.y += 1
				if calculateDistance(head.location.x, head.location.y, tail.location.x, tail.location.y) {
					tail.location.x = oldHeadX
					tail.location.y = oldHeadY
					tail.visitedCoords = append(tail.visitedCoords, Coordinate{oldHeadX, oldHeadY})
				}

			case "D":
				head.location.y -= 1
				if calculateDistance(head.location.x, head.location.y, tail.location.x, tail.location.y) {
					tail.location.x = oldHeadX
					tail.location.y = oldHeadY
					tail.visitedCoords = append(tail.visitedCoords, Coordinate{oldHeadX, oldHeadY})
				}

			case "L":
				head.location.x -= 1
				if calculateDistance(head.location.x, head.location.y, tail.location.x, tail.location.y) {
					tail.location.x = oldHeadX
					tail.location.y = oldHeadY
					tail.visitedCoords = append(tail.visitedCoords, Coordinate{oldHeadX, oldHeadY})
				}

			case "R":
				head.location.x += 1
				if calculateDistance(head.location.x, head.location.y, tail.location.x, tail.location.y) {
					tail.location.x = oldHeadX
					tail.location.y = oldHeadY
					tail.visitedCoords = append(tail.visitedCoords, Coordinate{oldHeadX, oldHeadY})
				}
			}
		}
	}

	mp := make(map[Coordinate]bool, 0)
	for _, coord := range tail.visitedCoords {
		mp[coord] = true
	}

	fmt.Println(len(mp) + 1)

}
