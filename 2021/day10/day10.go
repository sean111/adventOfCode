package main

import (
	"bufio"
	"log"
	"os"
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
	p1Total := 0
	data, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		var badChar rune = 0
		line := []rune(scanner.Text())
		log.Printf("%#v", line)
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
			log.Printf("%s = %d", string(badChar), value)
			p1Total += value
		}
	}

	log.Printf("Part 1 Answer: %d", p1Total)
}

func getRuneValue(delim rune, runeMap map[rune]RuneInfo) int {
	for _, info := range runeMap {
		if info.closing == delim {
			return info.value
		}
	}
	return 0
}
