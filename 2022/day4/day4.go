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

func (self Assignment) Overlaps(target Assignment) bool {
	if self.end <= target.end && self.end >= target.start {
		return true
	}
	return false
}

func main() {
	data, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}
	defer data.Close()

	scanner := bufio.NewScanner(data)
	var total int
	var p2Total int
	for scanner.Scan() {
		assignments := ParseInput(scanner.Text())
		if assignments[0].Contains(assignments[1]) || assignments[1].Contains(assignments[0]) {
			total++
		}

		if assignments[0].Overlaps(assignments[1]) || assignments[1].Overlaps(assignments[0]) {
			p2Total++
		}
	}
	fmt.Printf("Total: %d\n", total)
	fmt.Printf("P2 Total: %d\n", p2Total)
}

func ParseInput(input string) []Assignment {
	r := regexp.MustCompile(`(\d+)-(\d+)`)
	match := r.FindAllStringSubmatch(input, -1)
	assignments := make([]Assignment, 0)
	for _, m := range match {
		firstPoint, _ := strconv.Atoi(m[1])
		secondPoint, _ := strconv.Atoi(m[2])

		if firstPoint > secondPoint {
			assignments = append(assignments, Assignment{start: secondPoint, end: firstPoint})
		} else {
			assignments = append(assignments, Assignment{start: firstPoint, end: secondPoint})
		}
	}
	return assignments
}
