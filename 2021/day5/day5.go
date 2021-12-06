package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	var vents [1000][1000]int
	var overlaps = 0
	data, err := os.Open("test.txt")
	if err != nil {
		log.Fatalln(err)
		return
	}
	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		entry := scanner.Text()
		r, err := regexp.Compile("\\d+")

		if err != nil {
			log.Fatalln(err)
			return
		}
		coords := r.FindAllString(entry, -1)
		log.Printf("(X1, Y1): %s,%s => (X2,Y2): %s,%s\n", coords[0], coords[1], coords[2], coords[3])
		x1, _ := strconv.Atoi(coords[0])
		y1, _ := strconv.Atoi(coords[1])
		x2, _ := strconv.Atoi(coords[2])
		y2, _ := strconv.Atoi(coords[3])
		start := 0
		end := 0
		if x1 == x2 {
			if y1 > y2 {
				start = y2
				end = y1
			} else {
				start = y1
				end = y2
			}
			for i := start; i <= end; i++ {
				vents[x1][i]++
				if vents[x1][i] == 2 {
					overlaps++
				}
			}
		} else if y1 == y2 {
			if x1 > x2 {
				start = x2
				end = x1
			} else {
				start = x1
				end = x2
			}
			for i := start; i <= end; i++ {
				vents[i][y1]++
				if vents[i][y1] == 2 {
					overlaps++
				}
			}
		}
	}
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			if vents[x][y] == 0 {
				fmt.Printf(". ")
			} else {
				fmt.Printf("%d ", vents[x][y])
			}
		}
		fmt.Println("")
	}

	log.Printf("Overlaps: %d", overlaps)
}
