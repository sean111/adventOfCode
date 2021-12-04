package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var p1Answer = 0
	var p2Answer = 0
	var lastEntry = 0
	var lastSum = 0
	var entryArray[]int
	data, err := os.Open("data.txt")	
	if err != nil {
		log.Fatal(err)
		return
	}
	defer data.Close()

	buffer := bufio.NewScanner(data)

	for buffer.Scan() {
		entry, _ := strconv.Atoi(buffer.Text())
		if entry > lastEntry && lastEntry != 0 {
			p1Answer++
		}

		entryArray = append(entryArray, entry)
		lastEntry = entry
		
		length := len(entryArray)
		if length > 2 {
			tmpSum := entryArray[length-3] + entryArray[length-2] + entryArray[length-1] //TODO: Finda  proper way to do this
			if lastSum > 0 && tmpSum > lastSum {
				p2Answer++
			}
			lastSum = tmpSum
		}		
	}

	err = buffer.Err()
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("Part 1 Answer: %d\n", p1Answer)
	fmt.Printf("Part 2 Answer: %d\n", p2Answer)
}
