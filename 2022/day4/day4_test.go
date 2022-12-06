package main

import (
	"fmt"
	"testing"
)

func TestAssignmentContains(t *testing.T) {
	var tests = []struct {
		assignment1 Assignment
		assignment2 Assignment
		result      bool
	}{
		{assignment1: Assignment{start: 1, end: 3}, assignment2: Assignment{start: 2, end: 3}, result: true},
		{assignment1: Assignment{start: 6, end: 9}, assignment2: Assignment{start: 1, end: 7}, result: false},
		{assignment1: Assignment{start: 10, end: 11}, assignment2: Assignment{start: 9, end: 22}, result: true},
		{assignment1: Assignment{start: 13, end: 11}, assignment2: Assignment{start: 10, end: 12}, result: true},
		{assignment1: Assignment{start: 18, end: 12}, assignment2: Assignment{start: 13, end: 11}, result: false},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v, %v", test.assignment1, test.assignment2), func(t *testing.T) {
			ans := false

			if test.assignment1.Contains(test.assignment2) || test.assignment2.Contains(test.assignment1) {
				ans = true
			}

			if ans != test.result {
				t.Errorf("got %t, want %t", ans, test.result)
			}
		})
	}
}

func TestParseInput(t *testing.T) {
	var tests = []struct {
		input       string
		assignment1 Assignment
		assignment2 Assignment
	}{
		{
			input:       "17-72,16-71",
			assignment1: Assignment{start: 17, end: 72},
			assignment2: Assignment{start: 16, end: 71},
		},
		{
			input:       "3-31,1-32",
			assignment1: Assignment{start: 3, end: 31},
			assignment2: Assignment{start: 1, end: 32},
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			ans := ParseInput(test.input)
			if ans[0].start != test.assignment1.start || ans[0].end != test.assignment1.end || ans[1].start != test.assignment2.start || ans[1].end != test.assignment2.end {
				t.Errorf("got %v, %v, want %v, %v", ans[0], ans[1], test.assignment1, test.assignment2)
			}
		})
	}
}
