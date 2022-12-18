package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Position struct {
	row    int
	column int
}

var Map [][]bool

type Rope struct {
	head Position
	tail Position
}

var rope Rope

func main() {
	rows := 10
	columns := 10
	rope = Rope{head: Position{row: 4}, tail: Position{row: 4}}
	// Build base map
	for r := 0; r < rows; r++ {
		var tmpRow []bool
		for c := 0; c < columns; c++ {
			tmpRow = append(tmpRow, false)
		}
		Map = append(Map, tmpRow)
	}
	Map[rope.tail.row][rope.tail.column] = true
	fmt.Printf("Rope: %v\n", rope)
	ShowMap(true)

	data, err := os.Open("test.txt")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		rawInput := scanner.Text()
		dir := []rune(rawInput)

		spaces, _ := strconv.Atoi(string(dir[2]))

		rope = rope.Move(dir[0], spaces)
		fmt.Println()
		fmt.Printf("Rope: %v\n", rope)
		ShowMap(true)
	}

}

func ShowMap(showRope bool) {
	for r := 0; r < len(Map); r++ {
		for c := 0; c < len(Map[0]); c++ {
			if rope.IsHeadAt(r, c) && showRope {
				fmt.Print("H ")
			} else if rope.IsTailAt(r, c) && showRope {
				fmt.Print("T ")
			} else if Map[r][c] {
				fmt.Print("# ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
}

func (r Rope) IsHeadAt(row int, column int) bool {
	if r.head.row == row && r.head.column == column {
		return true
	}
	return false
}

func (r Rope) IsTailAt(row int, column int) bool {
	if r.tail.row == row && r.tail.column == column {
		return true
	}
	return false
}

func (r Rope) Move(direction rune, spaces int) Rope {
	log.Printf("%c -> %d", direction, spaces)
	//tmpPos := r.head
	for moved := 0; moved < spaces; moved++ {
		// Move head
		switch direction {
		case 'U':
			r.head.row--
		case 'D':
			r.head.row++
		case 'L':
			r.head.column--
		case 'R':
			r.head.column++
		}

		// Move Tail

	}

	return r
}
