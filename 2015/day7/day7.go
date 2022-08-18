package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	data, err := os.Open("test.txt")
	if err != nil {
		log.Fatalln(err)
	}
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {

	}
}
