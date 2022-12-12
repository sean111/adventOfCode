package main

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"log"
	"os"
	"strconv"
)

type Tree struct {
	height  int
	visible bool
	checked bool
}

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
			chkd := false
			if row == 0 || column == 0 || column == len(line)-1 {
				vis = true
				chkd = true
			}
			//Map[row][column] = Tree{height: tmp, visible: vis, checked: false}
			tmpRow = append(tmpRow, Tree{height: tmp, visible: vis, checked: chkd})
		}
		Map = append(Map, tmpRow)
		//column = 0
		row++
	}

	//Go back through and set last row to visible
	for column := 1; column < len(Map[row-1])-1; column++ {
		Map[row-1][column].visible = true
		Map[row-1][column].checked = true
	}

	fmt.Println("Initial Map")
	ShowMap()
	IsTreeVisible(1, 1)
	fmt.Println("Final Map")
	ShowMap()
}

func IsTreeVisible(row int, column int) bool {
	tree := &Map[row][column]
	log.Printf("Tree: %v\n", tree)

	var lVisible bool = false
	var rVisible bool = false
	var uVisible bool = false
	var dVisible bool = false

	// Check Left
	if column > 0 {
		lVisible = Map[row][column-1].visible
		lHeight := Map[row][column-1].height
		if Map[row][column-1].checked != true {
			lVisible = IsTreeVisible(row, column-1)
		}

		if lVisible == true && lHeight < tree.height {
			Map[row][column].visible = true
			tree.visible = true
		}
		Map[row][column].checked = true
	}

	// Check right
	if column < len(Map[row]) {
		rVisible = Map[row][column+1].visible
		rHeight := Map[row][column+1].height

		if Map[row][column+1].checked != true {
			rVisible = IsTreeVisible(row, column+1)
		}

		if rVisible == true && rHeight < tree.height {
			Map[row][column].visible = true
		}
		Map[row][column].checked = true
	}

	//Check Up
	if row > 0 {
		uVisible = Map[row-1][column].visible
		uHeight := Map[row-1][column].height

		if Map[row-1][column].checked != true {
			uVisible = IsTreeVisible(row-1, column)
		}

		if uVisible == true && uHeight < tree.height {
			Map[row][column].visible = true
		}
		Map[row][column].checked = true
	}

	if row < len(Map) {
		dVisible = Map[row+1][column].visible
		dHeight := Map[row+1][column].height

		if Map[row+1][column].checked != true {
			dVisible = IsTreeVisible(row+1, column)
		}

		if dVisible == true && dHeight < tree.height {
			Map[row][column].visible = true
		}
		Map[row][column].checked = true
	}

	if lVisible || rVisible || uVisible || dVisible {
		return true
	}
	return false
}

func ShowMap() {
	for row := 0; row < len(Map); row++ {
		for column := 0; column < len(Map[row]); column++ {
			tree := Map[row][column]
			//fmt.Printf("%d [%t,%t]\t", Map[row][column].height, Map[row][column].visible, Map[row][column].checked)
			outColor := color.New(color.FgGreen).SprintfFunc()
			if !tree.visible {
				outColor = color.New(color.FgRed).SprintfFunc()
			}
			tmp := strconv.Itoa(tree.height)
			fmt.Printf("%s\t", outColor(tmp))

		}
		fmt.Println()
	}
}
