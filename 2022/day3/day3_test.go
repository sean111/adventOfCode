package main

import (
	"fmt"
	"testing"
)

var Tests = []struct {
	input     string
	want      Rucksack
	duplicate rune
}{
	{"abaccb", Rucksack{Compartment1: []rune("aba"), Compartment2: []rune("ccb")}, 'b'},
	{"vJrwpWtwJgWrhcsFMMfFFhFp", Rucksack{Compartment1: []rune("vJrwpWtwJgWr"), Compartment2: []rune("hcsFMMfFFhFp")}, 'p'},
	{"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", Rucksack{Compartment1: []rune("jqHRNqRjqzjGDLGL"), Compartment2: []rune("rsFMfFZSrLrFZsSL")}, 'L'},
}

func TestParseContainer(t *testing.T) {
	for _, test := range Tests {
		testName := fmt.Sprintf("%s", test.input)
		t.Run(testName, func(t *testing.T) {
			ans := ParseContainer(test.input)
			if string(ans.Compartment1) != string(test.want.Compartment1) || string(ans.Compartment2) != string(test.want.Compartment2) {
				t.Errorf("got %v, want %v", ans, test.want)
			}
		})
	}
}

func TestFindDuplicates(t *testing.T) {
	for _, test := range Tests {
		sack := ParseContainer(test.input)
		testName := fmt.Sprintf("%v", sack)
		t.Run(testName, func(t *testing.T) {
			ans := FindDuplicates(sack)
			if test.duplicate != ans[0] {
				t.Errorf("got %c, want %c", ans, test.duplicate)
			}
		})
	}
}

func TestGetItemPriority(t *testing.T) {
	var tests = []struct {
		item     rune
		priority int
	}{
		{'A', 27},
		{'a', 1},
		{'p', 16},
		{'P', 42},
		{'s', 19},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("%c", test.item)
		t.Run(testName, func(t *testing.T) {
			ans := GetItemPriority(test.item)
			if ans != test.priority {
				t.Errorf("got %d, want %d", ans, test.priority)
			}
		})
	}
}
