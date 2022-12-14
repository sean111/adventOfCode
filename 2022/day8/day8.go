package main

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"os"
	"strconv"
)

type Tree struct {
	height      int
	visible     bool
	visibleFrom rune // For debugging
	score       int
}

var MaxScore int

var Map [][]Tree

func main() {
	data, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(data)

	var row int

	for scanner.Scan() {
		line := scanner.Text()
		var tmpRow []Tree
		for column, c := range line {
			tmp, _ := strconv.Atoi(string(c))
			//fmt.Printf("Row: %d, Column: %d, Value: %d,\n", row, column, tmp)
			vis := false
			if row == 0 || column == 0 || column == len(line)-1 {
				vis = true
			}
			tmpRow = append(tmpRow, Tree{height: tmp, visible: vis, visibleFrom: ' '})
		}
		Map = append(Map, tmpRow)
		//column = 0
		row++
	}

	//Go back through and set last row to visible
	for column := 1; column < len(Map[row-1])-1; column++ {
		Tree := &Map[row-1][column]
		Tree.visible = true
	}
	for r := 0; r <= len(Map)-1; r++ {
		for c := 0; c <= len(Map[1])-1; c++ {
			CheckTree(r, c)
		}
	}
	// CheckTree(4, 3)
	//fmt.Println("Final Map")
	//ShowMap()

	vTrees := 0
	for r := 0; r <= len(Map)-1; r++ {
		for c := 0; c <= len(Map[r])-1; c++ {
			if Map[r][c].visible == true {
				vTrees++
			}
		}
	}

	fmt.Printf("Real Visible Trees: %d\n", vTrees)
	fmt.Printf("Max Score: %d\n", MaxScore)
}

func CheckTree(row int, column int) {
	//log.Println()
	var foundTallerTree bool
	var tmpScore int
	tree := &Map[row][column]

	// Check Up
	foundTallerTree = false
	tmpScore = 0
	for i := row - 1; i >= 0; i-- {
		tmpScore++
		target := Map[i][column]
		if target.height >= tree.height {
			foundTallerTree = true
			break
		}
	}
	if !foundTallerTree {
		tree.visibleFrom = 'U'
		tree.visible = true
	}
	tree.score += tmpScore

	// Check Right
	foundTallerTree = false
	tmpScore = 0
	for i := column + 1; i < len(Map[row]); i++ {
		//log.Printf("- %d, %d ::", row, i)
		tmpScore++
		target := Map[row][i]
		if target.height >= tree.height {
			foundTallerTree = true
			break
		}
	}
	if !foundTallerTree {
		tree.visibleFrom = 'R'
		tree.visible = true
	}
	tree.score *= tmpScore

	// Check Left
	foundTallerTree = false
	tmpScore = 0
	for i := column - 1; i >= 0; i-- {
		tmpScore++
		target := Map[row][i]
		if target.height >= tree.height {
			foundTallerTree = true
			break
		}
	}
	if !foundTallerTree {
		tree.visibleFrom = 'L'
		tree.visible = true
	}
	tree.score *= tmpScore

	// Check Down
	foundTallerTree = false
	tmpScore = 0
	for i := row + 1; i < len(Map); i++ {
		tmpScore++
		target := Map[i][column]
		if target.height >= tree.height {
			foundTallerTree = true
			break
		}
	}
	if !foundTallerTree {
		tree.visibleFrom = 'D'
		tree.visible = true
	}
	tree.score *= tmpScore
	if tree.score > MaxScore {
		MaxScore = tree.score
	}
}

func ShowMap() {
	for row := 0; row < len(Map); row++ {
		for column := 0; column < len(Map[row]); column++ {
			tree := Map[row][column]
			outColor := color.New(color.FgGreen).SprintfFunc()
			if !tree.visible {
				outColor = color.New(color.FgRed).SprintfFunc()
			}
			tmp := strconv.Itoa(tree.height)
			fmt.Printf("%s[%d] ", outColor(tmp), tree.score)

		}
		fmt.Println()
	}
}
