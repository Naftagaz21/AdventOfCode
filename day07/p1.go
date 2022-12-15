package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func (d *dir) size() int {
	result := 0
	for _, file := range d.files {
		result += file.size
	}

	for _, dir := range d.dirs {
		result += dir.size()
	}
	return result
}

func (d *dir) forEach(f func(d *dir)) {
	for _, dir := range d.dirs {
		dir.forEach(f)
	}
	f(d)
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
			var file file
			fileSize, _ := strconv.Atoi(lines[0])
			file.name = lines[1]
			file.size = fileSize
			current.files = append(current.files, file)
		}
	}

	total := 0
	unused_space := 70000000 - root.size()
	min_space := 99999999
	root.forEach(func(d *dir) {
		size := d.size()
		if size <= 100000 {
			total += size
		}
		if size+unused_space >= 30000000 && size < min_space {
			min_space = size
		}
	})
	fmt.Println(total, min_space)
}
