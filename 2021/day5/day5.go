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
	data, err := os.Open("data.txt")
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

		vents[x1][y1]++
		if vents[x1][y1] == 2 {
			overlaps++
		}
		for {
			if x1 != x2 {
				if x1 > x2 {
					x1--
				} else {
					x1++
				}
			}
			if y1 != y2 {
				if y1 > y2 {
					y1--
				} else {
					y1++
				}
			}
			vents[x1][y1]++
			if vents[x1][y1] == 2 {
				overlaps++
			}

			if x1 == x2 && y1 == y2 {
				break
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
