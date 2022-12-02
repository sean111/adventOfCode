package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	data, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}
	defer data.Close()
	buffer := bufio.NewScanner(data)

	caloriesPerElf := make([]int, 0)
	currentSum := 0
	for buffer.Scan() {
		input := buffer.Text()
		if input == "" {
			caloriesPerElf = append(caloriesPerElf, currentSum)
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
	caloriesPerElf = append(caloriesPerElf, currentSum)

	sort.Ints(caloriesPerElf)
	sl := len(caloriesPerElf)

	p2Answer := caloriesPerElf[sl-1] + caloriesPerElf[sl-2] + caloriesPerElf[sl-3]

	fmt.Printf("P1 Answer: %d\n", caloriesPerElf[sl-1])
	fmt.Printf("P2 Answer: %d\n", p2Answer)
}
