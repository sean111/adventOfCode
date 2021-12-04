package main

import (
	"bufio"
	"os"
	"log"
	"strings"
	"strconv"
)

type Position struct {
	x int //Horizontal
	y int //Depth
}

type AdvPosition struct {
	x int
	y int
	aim int
}

func main() {	
	position := Position{x: 0, y: 0}
	advPosition := AdvPosition{x: 0, y: 0, aim: 0}
	data, err := os.Open("data.txt")	
	if err != nil {
		log.Fatal(err)
		return
	}
	defer data.Close()
	buffer := bufio.NewScanner(data)
	for buffer.Scan() {
		//TODO: use regex
		commands := strings.Split(buffer.Text(), " ")
		movementAmount, _ := strconv.Atoi(commands[1])
		switch commands[0] {
		case "forward":
			position.x += movementAmount
			advPosition.x += movementAmount
			advPosition.y += movementAmount * advPosition.aim
		case "up":
			position.y -= movementAmount
			advPosition.aim -= movementAmount
		case "down":
			position.y += movementAmount 
			advPosition.aim += movementAmount
		}		
	}

	log.Printf("Position %d,%d\n", position.x, position.y)
	p1Final := position.x * position.y
	p2Final := advPosition.x * advPosition.y

	log.Printf("Part 1: %d\n", p1Final)
	log.Printf("Part 2: %d\n", p2Final)
}
