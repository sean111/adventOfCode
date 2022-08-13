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

var santaCoords Coords
var roboSantaCoords Coords
var coords Coords

var moover int8 = 0 // 0 = Santa, 1 = RoboSanta

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
		if moover == 0 {
			coords = santaCoords
		} else {
			coords = roboSantaCoords
		}

		switch string(direction) {
		case ">":
			coords = move(coords.x+1, coords.y)
		case "<":
			coords = move(coords.x-1, coords.y)
		case "^":
			coords = move(coords.x, coords.y+1)
		case "v":
			coords = move(coords.x, coords.y-1)
		}

		if moover == 0 {
			moover = 1
			santaCoords = coords
		} else {
			moover = 0
			roboSantaCoords = coords
		}
	}
	log.Printf("Houses: %d", houses)
}

func move(x int, y int) Coords {
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
	return Coords{x: x, y: y}
}
