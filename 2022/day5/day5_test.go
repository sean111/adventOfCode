package main

import (
	"fmt"
	"log"
	"strings"
	"testing"
)

//var TestMap = map[int] []string {
//	1: {"Z", "N"},
//	2: {"D", "C", "M"},
//	3: {"P"},
//}

var TestMap = Cargo{
	columns: []Stack{
		{
			boxes: []string{
				"Z", "N",
			},
		},
		{
			boxes: []string{
				"M", "C", "D",
			},
		},
		{
			boxes: []string{
				"P",
			},
		},
	},
}

func TestCargo_MovePos(t *testing.T) {
	cargo := TestMap

	var tests = []struct {
		count  int
		from   int
		to     int
		result string
	}{
		{
			count:  1,
			from:   2,
			to:     1,
			result: "ZNDMCP",
		},
		{
			count:  3,
			from:   1,
			to:     3,
			result: "MCPDNZ",
		},
	}

	log.Printf("Test: %v\n", cargo)

	for _, test := range tests {
		t.Run(fmt.Sprintf("%d, %d, %d", test.count, test.from, test.to), func(t *testing.T) {
			cargo.MovePos(test.count, test.from, test.to)
			ans := cargoToString(cargo)
			if ans != test.result {
				t.Errorf("got %s, want %s", ans, test.result)
			}
		})
	}
}

func TestCargo_GetTopCrates(t *testing.T) {
	cargo := TestMap
	want := "NDP"
	ans := cargo.GetTopCrates()

	if ans != want {
		t.Errorf("got %s, want %s", ans, want)
	}

}

func cargoToString(cargo Cargo) string {
	var tmp []string
	for _, column := range cargo.columns {
		for _, box := range column.boxes {
			tmp = append(tmp, box)
		}
	}
	return strings.Join(tmp[:], "")
}
