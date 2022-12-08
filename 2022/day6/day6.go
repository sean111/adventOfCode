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
	input := string(data)
	result := FindSignal(input, 4)
	result2 := FindSignal(input, 14)

	fmt.Printf("Character Pos: %d\n", result)
	fmt.Printf("P2 Character Pos: %d", result2)
}

func FindSignal(input string, count int) int {
	//log.Printf("Input: %s\n", input)
	chars := []rune(input)
	//log.Printf("Chars: %v\n", chars)
	charLen := len(chars)
	for i := count; i < charLen; i++ {
		test := chars[i-count : i]
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
