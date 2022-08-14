package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var badStrings = [4]string{"ab", "cd", "pq", "xy"}
var vowels = [5]string{"a", "e", "i", "o", "u"}

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
		log.Printf("String: %s", inputString)
		vowelCount := getVowlCount(inputString)
		log.Printf("\tVowels: %d", vowelCount)
		if vowelCount < 3 {
			log.Printf("\tFailed vowel count (%d)\n", vowelCount)
			continue
		}

		if !checkForDoubleLetter(inputString) {
			log.Printf("\tFailed double letter check\n")
			continue
		}

		if checkForBadStrings(inputString) {
			log.Printf("\tFailed bad string check\n")
			continue
		}

		goodStrings++
	}
	fmt.Printf("Good Strings: %d\n", goodStrings)
}

func getVowlCount(target string) int {
	var total int = 0
	for _, vowel := range vowels {
		total += strings.Count(target, vowel)
	}
	return total
}

func checkForDoubleLetter(target string) bool {
	for x := 'a'; x <= 'z'; x++ {
		if strings.Contains(target, string(x)+string(x)) {
			return true
		}
	}
	return false
}

func checkForBadStrings(target string) bool {
	for _, badString := range badStrings {
		if strings.Contains(target, badString) {
			return true
		}
	}
	return false
}
