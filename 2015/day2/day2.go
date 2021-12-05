package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	totalSurface := 0
	totalRibbon := 0
	data, err := os.Open("test	.txt")
	if err != nil {
		log.Fatal(err)
		return
	}

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		line := scanner.Text()
		dims := strings.Split(line, "x")
		length, _ := strconv.Atoi(dims[0])
		width, _ := strconv.Atoi(dims[1])
		height, _ := strconv.Atoi(dims[2])
		lwVal := length * width
		whVal := width * height
		hlVal := height * length
		surface := 2*lwVal + 2*whVal + 2*hlVal + min(lwVal, whVal, hlVal)
		ribbon := 0
		totalRibbon += ribbon
		log.Printf("%s => %d (surface) => %d (ribbon)\n", line, surface, ribbon)
		totalSurface += surface
	}
	log.Printf("Total Surface: %d | Total Ribbon: %d", totalSurface, totalRibbon)
}

func min(values ...int) int {
	var min int
	for i, val := range values {
		if i == 0 || min > val {
			min = val
		}
	}
	return min
}
