package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var goodStrings int = 0
	data, err := os.Open("data.txt")

	if err != nil {
		log.Fatal(err)
		return
	}
	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		inputString := scanner.Text()
		log.Printf("String: %s\n", inputString)

		if !checkForPairs(inputString) {
			log.Printf("\tFailed pair check\n")
			continue
		}

		if !checkForPattern(inputString) {
			log.Printf("\tFailed pattern check\n")
			continue
		}

		goodStrings++
	}
	fmt.Printf("Good Strings: %d\n", goodStrings)
}

func checkForPairs(target string) bool {
	targetLen := len([]rune(target)) - 2
	for x := 0; x <= targetLen; x++ {
		//log.Printf("\t\tsubstr: %s\n", target[x:x+2])
		fCount := strings.Count(target, target[x:x+2])
		if fCount > 1 {
			//log.Printf("\t\t%s => %d", target[x:x+2], fCount)
			return true
		}
	}
	return false
}

func checkForPattern(target string) bool {
	targetLen := len([]rune(target)) - 3
	for x := 0; x <= targetLen; x++ {
		if target[x] == target[x+2] {
			return true
		}
	}
	return false
}
