package main

import (
	"io/ioutil"
	"strings"
)

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

func main() {

	file, _ := ioutil.ReadFile("input.txt")
	input := string(file)

	parts := strings.Split(input, "\n\n")

	stackPart := parts[0]
	stacks := strings.Split(stackPart, "\n")

	nrOfStacks := (len(stacks[len(stacks)-1])-1)/4 + 1
	stackStore := make([]Stack, nrOfStacks)

	//fmt.Println(stacks)

}
