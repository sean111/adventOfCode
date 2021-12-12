package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.Open("data.txt")
	if err != nil {
		log.Fatalln(err)
		return
	}
	scanner := bufio.NewScanner(data)
	p1Total := 0
	p2Total := 0
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "|")
		oneVal, fourVal := getKeyValues(line[0])
		resultEntries := strings.Split(strings.TrimLeft(line[1], " "), " ")
		var output string
		for _, entry := range resultEntries {
			strlen := len(entry)
			n1Common := commonChars(oneVal, entry)
			n4Common := commonChars(fourVal, entry)
			switch strlen {
			case 2:
				p1Total++
				output += "1"
			case 3:
				p1Total++
				output += "7"
			case 4:
				p1Total++
				output += "4"
			case 7:
				p1Total++
				output += "8"
			case 5:
				if n1Common == 1 && n4Common == 2 {
					output += "2"
				} else if n1Common == 2 && n4Common == 3 {
					output += "3"
				} else {
					output += "5"
				}
			case 6:
				if n1Common == 2 && n4Common == 4 {
					output += "9"
				} else if n1Common == 1 && n4Common == 3 {
					output += "6"
				} else {
					output += "0"
				}
			}
		}
		log.Printf("Line Output: %s \n", output)
		outputTmp, _ := strconv.Atoi(output)
		p2Total += outputTmp
	}
	log.Printf("Part 1 Total: %d\n", p1Total)
	log.Printf("Part 2 Total: %d\n", p2Total)
}

func getKeyValues(line string) (string, string) {
	oneString, fourString := "", ""

	entries := strings.Split(line, " ")

	for _, entry := range entries {
		if len(entry) == 2 {
			oneString = entry
		} else if len(entry) == 4 {
			fourString = entry
		}
	}
	return oneString, fourString
}

func commonChars(a, b string) int {
	characters := ""
	for _, character := range a {
		if strings.ContainsRune(b, character) {
			characters += string(character)
		}
	}
	return len(characters)
}
