package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

type RuneInfo struct {
	closing rune
	value   int
}

func main() {
	runeMap := map[rune]RuneInfo{
		'(': {closing: ')', value: 3},
		'[': {closing: ']', value: 57},
		'{': {closing: '}', value: 1197},
		'<': {closing: '>', value: 25137},
	}
	corruptRuneMap := map[rune]int{
		'(': 1,
		'[': 2,
		'{': 3,
		'<': 4,
	}
	p1Total := 0
	var p2Totals []int
	data, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		var badChar rune = 0
		line := []rune(scanner.Text())
		// log.Printf("%#v", line)
		valid := true
		dBuffer := make([]rune, 0, len(line))

		for _, char := range line {
			if _, ok := runeMap[char]; ok {
				dBuffer = append(dBuffer, char)
			} else if len(dBuffer) == 0 {
				valid = false
				badChar = char
				break
			} else if runeMap[dBuffer[len(dBuffer)-1]].closing != char {
				valid = false
				badChar = char
				break
			} else {
				dBuffer = dBuffer[:len(dBuffer)-1] //remove the last char
			}
		}

		if !valid {
			value := getRuneValue(badChar, runeMap)
			p1Total += value
			dBuffer = dBuffer[:0]
		}

		if len(dBuffer) > 0 {
			lineTotal := 0
			for i := len(dBuffer) - 1; i >= 0; i-- {
				lineTotal = (lineTotal * 5) + getOpenRuneValue(dBuffer[i], corruptRuneMap)
			}
			p2Totals = append(p2Totals, lineTotal)
		}
	}

	log.Printf("Part 1 Answer: %d", p1Total)
	fmt.Println("PART 2 DEBUG")
	sort.Ints(p2Totals)
	log.Printf("Part 2 Answer: %d", p2Totals[((len(p2Totals)-1)/2)])
}

func getOpenRuneValue(delim rune, badRuneMap map[rune]int) int {
	_, found := badRuneMap[delim]
	if found {
		return badRuneMap[delim]
	}
	return 0
}

func getRuneValue(delim rune, runeMap map[rune]RuneInfo) int {
	for _, info := range runeMap {
		if info.closing == delim {
			return info.value
		}
	}
	return 0
}
