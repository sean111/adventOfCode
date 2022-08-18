package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Coord struct {
	x int
	y int
}

type Operation struct {
	command string
	start   Coord
	end     Coord
}

var grid [1000][1000]int

var lightsLit int

func main() {
	data, err := os.Open("data.txt")
	if err != nil {
		log.Fatalln(err)
	}
	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		command := scanner.Text()
		operation := commandParser(command)
		for x := operation.start.x; x <= operation.end.x; x++ {
			for y := operation.start.y; y <= operation.end.y; y++ {
				switch operation.command {
				case "toggle":
					if grid[x][y] == 1 {
						lightsLit--
						grid[x][y] = 0
					} else {
						lightsLit++
						grid[x][y] = 1
					}
				case "on":
					if grid[x][y] == 0 {
						lightsLit++
					}
					grid[x][y] = 1
				case "off":
					if grid[x][y] == 1 {
						lightsLit--
					}
					grid[x][y] = 0
				}
			}
		}
	}
	//outputGrid()
	fmt.Printf("Lights Lit: %d\n", lightsLit)
	check := totalLightsLit()
	fmt.Printf("Lights Lit Check: %d\n", check)
}

func commandParser(line string) Operation {
	log.Printf("Line: %s\n", line)
	var command string
	var start Coord
	var end Coord

	var startPos int

	if string(line[1]) == "o" {
		command = "toggle"
		startPos = 1
	} else {
		if string(line[6]) == "n" {
			command = "on"
		} else {
			command = "off"
		}
		startPos = 2
	}

	tmpLine := strings.Fields(line)
	tmpStart := strings.Split(tmpLine[startPos], ",")
	start.x, _ = strconv.Atoi(tmpStart[0])
	start.y, _ = strconv.Atoi(tmpStart[1])
	strLen := len(tmpLine)
	tmpEnd := strings.Split(tmpLine[strLen-1], ",")
	end.x, _ = strconv.Atoi(tmpEnd[0])
	end.y, _ = strconv.Atoi(tmpEnd[1])

	log.Printf("\tCommand: %s, start: %v, end: %v\n", command, start, end)
	return Operation{
		command: command,
		start:   start,
		end:     end,
	}
}

func outputGrid() {
	output, _ := os.Create("grid.txt")
	writer := bufio.NewWriter(output)
	xLen := len(grid)
	yLen := len(grid[0])
	for y := 0; y < yLen; y++ {
		outputString := ""
		for x := 0; x < xLen; x++ {
			if x == xLen-1 {
				outputString += fmt.Sprintf("\t%d (%d,%d)\n", grid[x][y], x, y)
			} else if x != 0 {
				outputString += fmt.Sprintf("\t%d (%d,%d)", grid[x][y], x, y)
			} else {
				outputString += fmt.Sprintf("%d (%d,%d)", grid[x][y], x, y)
			}
		}
		writer.WriteString(outputString)
	}
	writer.Flush()
}

func totalLightsLit() int {
	lights := 0
	xLen := len(grid)
	yLen := len(grid[0])
	for x := 0; x < xLen; x++ {
		for y := 0; y < yLen; y++ {
			if grid[x][y] == 1 {
				lights++
			}
		}
	}
	return lights
}
