package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Value struct {
	value  int
	marked bool
}

type Board [5][5]Value

var boardTotals []int

func main() {
	var boards []Board
	var drawnNumbers []string
	var boardCompleted []bool
	var boardTotal int
	tmpBoard := Board{}
	rowCount := 0
	line := 1
	data, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		value := scanner.Text()
		if line == 1 {
			drawnNumbers = strings.Split(value, ",")
		} else {
			if len(value) == 0 {
				if line > 3 {
					boards = append(boards, tmpBoard)
					boardTotals = append(boardTotals, boardTotal)
					boardCompleted = append(boardCompleted, false)
				}
				boardTotal = 0
				tmpBoard = Board{}
				rowCount = 0
			} else {
				r, _ := regexp.Compile(`(\d+)`)
				rowValues := r.FindAllString(value, -1)
				for columnCount, rowValue := range rowValues {
					number, _ := strconv.Atoi(rowValue)
					boardTotal += number
					tmpBoard[rowCount][columnCount] = Value{value: number, marked: false}
				}
				rowCount++
			}
		}
		line++
	}
	boards = append(boards, tmpBoard) //Add last board onto the array
	boardTotals = append(boardTotals, boardTotal)
	boardCompleted = append(boardCompleted, false)
	log.Printf("Board Count: %d\n", len(boards))

	for _, drawnNumber := range drawnNumbers {
		number, _ := strconv.Atoi(drawnNumber)

		for i, board := range boards {
			if boardCompleted[i] {
				continue
			}
			resArray, found, bingo := checkBoard(board, number)
			if found {
				boards[i] = resArray
				boardTotals[i] -= number
				//Check if board is completed
				if bingo {
					log.Println("Bingo!!!!!!")
					debugBoard(boards[i])
					log.Printf("Board Total: %d", boardTotals[i])
					product := boardTotals[i] * number
					log.Printf("Board Product: %d", product)
					boardCompleted[i] = true // I'd rather remove completed boards but the way I did it was causing index out of range issues
				}
			}
		}
	}
}

func checkBoard(board Board, value int) (Board, bool, bool) {
	for x, row := range board {
		for y, entry := range row {
			if entry.value == value {
				board[x][y].marked = true
				return board, true, checkBingo(board, x, y)
			}
		}
	}
	return board, false, false
}

func checkBingo(board Board, row int, column int) bool {
	markedSum := 0
	// Check row
	for x := 0; x < 5; x++ {
		if board[x][column].marked {
			markedSum++
		}
	}

	if markedSum == 5 {
		return true
	}

	//Check column
	markedSum = 0
	for y := 0; y < 5; y++ {
		if board[row][y].marked {
			markedSum++
		}
	}

	if markedSum == 5 {
		return true
	}

	return false
}

func debugBoard(board Board) {
	for _, row := range board {
		for _, entry := range row {
			var output string
			if entry.marked {
				output = fmt.Sprintf("\033[1;34m%d\033[0m", entry.value)
			} else {
				output = strconv.Itoa(entry.value)
			}
			fmt.Printf("%s\t", output)
		}
		fmt.Println("")
	}
	fmt.Println("")
}
