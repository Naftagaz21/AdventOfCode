package main

//https://github.com/pin2t/aoc2022/blob/main/07.go
import (
	"bufio"
	"os"
	"strings"
)

type file struct {
	name string
	size int
}

type dir struct {
	name   string
	parent *dir
	files  []file
	dirs   []*dir
}

func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()

	root := dir{name: "/", files: []file{}, dirs: []*dir{}}
	current := &root
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		if line[:4] == "$ ls" {
			continue
		} else if line[:5] == "$ cd " {
			switch line[5:] {
			case "/":
				current = &root
			case "..":
				current = current.parent
			default:
				for _, dir := range current.dirs {
					if dir.name == line[5:] {
						current = dir
					}
				}
			}
		} else if line[:3] == "dir" {
			current.dirs = append(current.dirs, &dir{name: line[4:], files: []file{}, parent: current, dirs: []*dir{}})

		} else {
			lines := strings.Fields(line)
			
		}

	}
}
