package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestRound(t *testing.T) {
	data, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		roundInput := strings.Split(scanner.Text(), " ")
		opponentHand := HandToShape[roundInput[0]]
		ourHand := HandToShape[roundInput[1]]
		score, _ := strconv.Atoi(roundInput[2])
		testname := fmt.Sprintf("Round(%s, %s)", opponentHand, ourHand)
		t.Run(testname, func(t *testing.T) {
			ans := Round(opponentHand, ourHand)
			if ans != score {
				t.Errorf("got %d, want %d", ans, score)
			}
		})
	}
}
