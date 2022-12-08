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
		count     int
		character int
	}{
		{
			input:     "bvwbjplbgvbhsrlpgdmjqwftvncz",
			count:     4,
			character: 5,
		},
		{
			input:     "nppdvjthqldpwncqszvftbrmjlhg",
			count:     4,
			character: 6,
		},
		{
			input:     "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			count:     4,
			character: 10,
		},
		{
			input:     "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			count:     4,
			character: 11,
		},
		{
			input:     "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			count:     14,
			character: 19,
		},
		{
			input:     "bvwbjplbgvbhsrlpgdmjqwftvncz",
			count:     14,
			character: 23,
		},
		{
			input:     "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			count:     14,
			character: 29,
		},
		{
			input:     "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			count:     14,
			character: 26,
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			ans := FindSignal(test.input, test.count)

			if ans != test.character {
				t.Errorf("got %d, want %d", ans, test.character)
			}
		})
	}
}
