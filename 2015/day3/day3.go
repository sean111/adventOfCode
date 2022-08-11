package main

import (
	"bufio"
	"log"
	"os"
)

var houses int
var houseMap = map[int]map[int]int{}

type Coords struct {
	x int
	y int
}

var coords Coords

func main() {
	data, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
		return
	}

	houses = 1
	houseMap[0] = map[int]int{
		0: 1,
	} //Initial delivery
	coords = Coords{x: 0, y: 0}

	reader := bufio.NewReader(data)
	for {
		direction, err := reader.ReadByte()
		if err != nil {
			if err.Error() == "EOF" {
				break
			} else {
				log.Fatalln(err)
			}
		}

		// log.Printf("%s", string(direction))
		switch string(direction) {
		case ">":
			move(coords.x+1, coords.y)
		case "<":
			move(coords.x-1, coords.y)
		case "^":
			move(coords.x, coords.y+1)
		case "v":
			move(coords.x, coords.y-1)
		}
	}
	log.Printf("Houses: %d", houses)
}

func move(x int, y int) {
	_, ok := houseMap[x][y]
	if !ok {
		if len(houseMap[x]) > 0 {
			houseMap[x][y] = 1
		} else {
			houseMap[x] = map[int]int{
				y: 1,
			}
		}
		houses++
	} else {
		houseMap[x][y]++
	}
	// log.Printf("%v\n", ok)
	// log.Printf("x: %d, y: %d, testVal: %d\n", x, y, testVal)
	// log.Printf("map: %v\n", houseMap)
	coords.x = x
	coords.y = y
}
