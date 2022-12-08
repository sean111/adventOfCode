package main

import (
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}
	// I don't feel like messing with []byte
	result := FindSignal(string(data))

	fmt.Printf("Character: %d\n", result)
}

func FindSignal(input string) int {
	//log.Printf("Input: %s\n", input)
	chars := []rune(input)
	//log.Printf("Chars: %v\n", chars)
	charLen := len(chars)
	for i := 4; i < charLen; i++ {
		test := chars[i-4 : i]
		//log.Printf("Test: %v\n", test)
		if !DupesFound(test) {
			return i
			break
		}
	}
	return 0
}

func DupesFound(chars []rune) bool {
	charLen := len(chars)
	for i := 0; i < charLen; i++ {
		for x := 0; x < charLen; x++ {
			if x != i && chars[i] == chars[x] {
				return true
			}
		}
	}
	return false
}
