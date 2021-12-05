package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	floor := 0
	data, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
		return
	}

	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		input := scanner.Text()
		for cnt, command := range []rune(input) {
			if command == '(' {
				floor++
			} else {
				floor--
			}

			if floor == -1 {
				log.Printf("In basement at command %d", cnt+1)
			}
		}
	}

	log.Printf("Final Floor: %d\n", floor)
}
