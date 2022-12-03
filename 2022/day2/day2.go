package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var shapeScores = map[string]int{
	"R": 1,
	"P": 2,
	"S": 3,
}

var HandToShape = map[string]string{
	"A": "R",
	"B": "P",
	"C": "S",
	"X": "R",
	"Y": "P",
	"Z": "S",
}

var winPoints = 6
var drawPoints = 3
var lossPoints = 0

func main() {
	data, err := os.Open("data.txt")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(data)

	total := 0

	for scanner.Scan() {
		roundInput := strings.Split(scanner.Text(), " ")
		total += Round(HandToShape[roundInput[0]], HandToShape[roundInput[1]])
	}
	fmt.Printf("Total: %d\n", total)
}

func Round(opponentHand string, ourHand string) int {
	// Get the hand we should play
	roundScore := shapeScores[ourHand]
	round := fmt.Sprintf("%s%s", opponentHand, ourHand)
	switch round {
	case "RS", "PR", "SP":
		roundScore += lossPoints
	case "SR", "RP", "PS":
		roundScore += winPoints
	default:
		roundScore += drawPoints
	}
	return roundScore
}
