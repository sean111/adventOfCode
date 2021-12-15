package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Cell struct {
	value  int
	lowest bool
}

func main() {
	var p1Result int
	var caves [][]Cell
	data, err := os.Open("test.txt")
	if err != nil {
		log.Fatalln(err)
		return
	}
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		var row []Cell
		line := scanner.Text()
		values := []rune(line)
		for _, val := range values {
			tmp, _ := strconv.Atoi(string(val))
			row = append(row, Cell{value: tmp, lowest: false})
		}
		caves = append(caves, row)
	}
	//debugCaves(caves)
	p1Result, caves = scanCaves(caves)
	//fmt.Println()
	debugCaves(caves)
	fmt.Printf("Part 1 Result: %d\n", p1Result)
}

func debugCaves(caves [][]Cell) {
	for _, row := range caves {
		for _, cell := range row {
			if cell.lowest {
				fmt.Printf("\033[1m%d\033[0m\t", cell.value)
			} else {
				fmt.Printf("%d\t", cell.value)
			}
		}
		fmt.Println("")
	}
}

func scanCaves(caves [][]Cell) (int, [][]Cell) {
	sum := 0
	for r, row := range caves {
		for c, cell := range row {
			uNum, lNum, dNum, rNum := 0, 0, 0, 0
			if r == 0 {
				uNum = 999
			} else {
				uNum = caves[r-1][c].value
			}

			if r == len(caves)-1 {
				dNum = 999
			} else {
				dNum = caves[r+1][c].value
			}

			if c == 0 {
				lNum = 999
			} else {
				lNum = caves[r][c-1].value
			}

			if c == len(row)-1 {
				rNum = 999
			} else {
				rNum = caves[r][c+1].value
			}

			if uNum > cell.value && lNum > cell.value && dNum > cell.value && rNum > cell.value {
				caves[r][c].lowest = true
				sum += cell.value + 1
			}
		}
	}
	return sum, caves
}
