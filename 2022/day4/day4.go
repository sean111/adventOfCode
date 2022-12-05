package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Assignment struct {
	start int
	end   int
}

func (self Assignment) Contains(target Assignment) bool {
	if self.start >= target.start && self.end <= target.end {
		return true
	}
	return false
}

func main() {
	data, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(data)
	var total int
	for scanner.Scan() {
		assignments := ParseInput(scanner.Text())
		if assignments[0].Contains(assignments[1]) || assignments[1].Contains(assignments[0]) {
			total++
		}
	}
	fmt.Printf("Total: %d\n", total)
}

func ParseInput(input string) []Assignment {
	r := regexp.MustCompile(`(\d)-(\d)`)
	match := r.FindAllStringSubmatch(input, -1)
	assignments := make([]Assignment, 0)
	for _, m := range match {
		startPoint, _ := strconv.Atoi(m[1])
		endPoint, _ := strconv.Atoi(m[2])
		assignments = append(assignments, Assignment{start: startPoint, end: endPoint})

	}
	return assignments
}
