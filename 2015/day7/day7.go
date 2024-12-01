package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var Signals = make(map[string]int)

func main() {
	data, err := os.Open("data.txt")
	if err != nil {
		log.Fatalln(err)
	}
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		ParseLine(scanner.Text())
	}
	fmt.Printf("Line A: %d\n ", Signals["a"])
}

func ParseLine(data string) {
	tmp := strings.Split(data, " ")
	switch {
	case tmp[0][0] == 'N':
		value := ^Signals[tmp[1]]
		fmt.Printf("^%d = %d", Signals[tmp[1]], value)
		fmt.Printf("Assigning to %s\n", tmp[3])
		Signals[tmp[3]] = value
		break
	case tmp[1][0] == 'A':
		fmt.Println("AND")
		value := Signals[tmp[0]] & Signals[tmp[2]]
		fmt.Printf("%d & %d = %d\n ", Signals[tmp[0]], Signals[tmp[2]], value)
		fmt.Printf("Assigning to %s\n", tmp[4])
		Signals[tmp[4]] = value
	case tmp[1][0] == '-':
		value, _ := strconv.Atoi(tmp[0])
		fmt.Printf("Assigning %d to %s\n", value, tmp[2])
		Signals[tmp[2]] = value
		break
	case tmp[1][0] == 'L':
		tmpVal, _ := strconv.Atoi(tmp[2])
		value := Signals[tmp[0]] << tmpVal
		fmt.Printf("%d << %d = %d\n ", Signals[tmp[0]], tmpVal, value)
		fmt.Printf("Assigning to %s\n", tmp[4])
		break
	case tmp[1][0] == 'R':
		tmpVal, _ := strconv.Atoi(tmp[2])
		value := Signals[tmp[0]] >> tmpVal
		fmt.Printf("%d >> %d = %d\n ", Signals[tmp[0]], tmpVal, value)
		fmt.Printf("Assigning to %s\n", tmp[4])
		Signals[tmp[4]] = value
		break
	case tmp[1][0] == 'O':
		value := Signals[tmp[0]] | Signals[tmp[2]]
		fmt.Printf("%d | %d = %d\n ", Signals[tmp[0]], Signals[tmp[2]], value)
		fmt.Printf("Assigning to %s\n", tmp[4])
		Signals[tmp[4]] = value
		break
	}
}
