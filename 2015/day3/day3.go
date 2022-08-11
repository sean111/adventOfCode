package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	data, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
		return
	}

	var houseMap [][]int
	houses := 1

	houseMap[0][0] = 1 //Initial delivery

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

		log.Printf("%s", string(direction))
	}
}
