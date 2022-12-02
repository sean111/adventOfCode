package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	data, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}
	defer data.Close()
	buffer := bufio.NewScanner(data)

	currentSum := 0
	maxSum := 0
	for buffer.Scan() {
		input := buffer.Text()
		if input == "" {
			if currentSum > maxSum {
				maxSum = currentSum
				fmt.Printf("New Leader!\n")
			}
			currentSum = 0
		} else {
			val, err := strconv.Atoi(input)
			if err != nil {
				panic(err)
			}
			currentSum += val
		}
	}

	// Account for end of the input being the max
	if currentSum > maxSum {
		maxSum = currentSum
	}

	fmt.Printf("Max Sum: %d\n", maxSum)
}
