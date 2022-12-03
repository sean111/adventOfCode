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

	p1Total := 0
	p2Total := 0

	for scanner.Scan() {
		roundInput := strings.Split(scanner.Text(), " ")
		p1Total += Round(HandToShape[roundInput[0]], HandToShape[roundInput[1]])
		p2Total += P2Round(HandToShape[roundInput[0]], roundInput[1])
	}
	fmt.Printf("Total: %d\n", p1Total)
	fmt.Printf("P2 Total: %d\n", p2Total)
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

func P2Round(opponentHand string, neededOutcome string) int {
	var roundScore int
	switch neededOutcome {
	case "X":
		// Loose
		roundScore = lossPoints + shapeScores[GetLoosingPlay(opponentHand)]
	case "Z":
		// Win
		roundScore = winPoints + shapeScores[GetWiningPlay(opponentHand)]
	default:
		// Draw
		roundScore = drawPoints + shapeScores[opponentHand]
	}
	return roundScore
}

func GetLoosingPlay(play string) string {
	var hand string
	switch play {
	case "R":
		hand = "S"
	case "P":
		hand = "R"
	case "S":
		hand = "P"
	}
	return hand
}

func GetWiningPlay(play string) string {
	var hand string
	switch play {
	case "R":
		hand = "P"
	case "P":
		hand = "S"
	case "S":
		hand = "R"
	}
	return hand
}
