package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
)

func main() {
	data, err := os.Open("test.txt")
	if err != nil {
		log.Fatalln(err)
		return
	}
	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		entry := scanner.Text()
		r, err := regexp.Compile("[(0-99)]")

		if err != nil {
			log.Fatalln(err)
			return
		}
		coords := r.FindAllString(entry, -1)
		log.Printf("(X1, Y1): %s,%s => (X2,Y2): %s,%s\n", coords[0], coords[1], coords[2], coords[3])
	}
}
