package main

import (
	"fmt"
	"testing"
)

func TestDupesFound(t *testing.T) {
	var tests = []struct {
		input  []rune
		result bool
	}{
		{
			input:  []rune("bvwb"),
			result: true,
		},
		{
			input:  []rune("abcd"),
			result: false,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.input), func(t *testing.T) {
			ans := DupesFound(test.input)

			if ans != test.result {
				t.Errorf("got %t, want %t", ans, test.result)
			}
		})
	}
}

func TestFindSignal(t *testing.T) {
	var tests = []struct {
		input     string
		character int
	}{
		{
			input:     "bvwbjplbgvbhsrlpgdmjqwftvncz",
			character: 5,
		},
		{
			input:     "nppdvjthqldpwncqszvftbrmjlhg",
			character: 6,
		},
		{
			input:     "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			character: 10,
		},
		{
			input:     "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			character: 11,
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			ans := FindSignal(test.input)

			if ans != test.character {
				t.Errorf("got %d, want %d", ans, test.character)
			}
		})
	}
}
