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
		testname := fmt.Sprintf("%s, %s", opponentHand, ourHand)
		t.Run(testname, func(t *testing.T) {
			ans := Round(opponentHand, ourHand)
			if ans != score {
				t.Errorf("got %d, want %d", ans, score)
			}
		})
	}
}

func TestP2Round(t *testing.T) {
	data, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		roundInput := strings.Split(scanner.Text(), " ")
		opponentHand := HandToShape[roundInput[0]]
		neededOutcome := roundInput[1]
		score, _ := strconv.Atoi(roundInput[3])
		testname := fmt.Sprintf("%s, %s", opponentHand, neededOutcome)
		t.Run(testname, func(t *testing.T) {
			ans := P2Round(opponentHand, neededOutcome)
			if ans != score {
				t.Errorf("got %d, want %d", ans, score)
			}
		})
	}
}

func TestGetLoosingPlay(t *testing.T) {
	var tests = []struct {
		play   string
		result string
	}{
		{"R", "S"},
		{"P", "R"},
		{"S", "P"},
	}

	for _, test := range tests {
		testname := fmt.Sprintf("%s", test.play)
		t.Run(testname, func(t *testing.T) {
			ans := GetLoosingPlay(test.play)
			if ans != test.result {
				t.Errorf("got %s, want %s", ans, test.result)
			}

		})
	}
}

func TestGetWiningPlay(t *testing.T) {
	var tests = []struct {
		play   string
		result string
	}{
		{"R", "P"},
		{"P", "S"},
		{"S", "R"},
	}

	for _, test := range tests {
		testname := fmt.Sprintf("%s", test.play)
		t.Run(testname, func(t *testing.T) {
			ans := GetWiningPlay(test.play)
			if ans != test.result {
				t.Errorf("got %s, want %s", ans, test.result)
			}

		})
	}
}
