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
	defer data.Close()

	writeFile, err := os.OpenFile("output.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)

	if err != nil {
		panic(err)
	}
	defer writeFile.Close()

	scanner := bufio.NewScanner(data)
	var total int
	for scanner.Scan() {
		assignments := ParseInput(scanner.Text())
		if assignments[0].Contains(assignments[1]) || assignments[1].Contains(assignments[0]) {
			_, err := writeFile.WriteString(fmt.Sprintf("%d - %d, %d - %d || [%s]\n", assignments[0].start, assignments[0].end, assignments[1].start, assignments[1].end, scanner.Text()))
			if err != nil {
				panic(err)
			}
			total++
		}
	}
	writeFile.Sync()
	fmt.Printf("Total: %d\n", total)
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
