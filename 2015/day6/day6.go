package main

import (
	"bufio"
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

var operations []Operation

func main() {
	data, err := os.Open("test.txt")
	if err != nil {
		log.Fatalln(err)
	}
	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		command := scanner.Text()
		commandParser(command)
	}
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

func show
