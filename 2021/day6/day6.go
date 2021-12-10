package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var age [9]int
	data, err := os.ReadFile("data.txt")
	if err != nil {
		log.Fatalln(err)
		return
	}
	initialState := strings.Split(strings.TrimSuffix(string(data), "\n"), ",")

	for _, val := range initialState {
		life, _ := strconv.Atoi(val)
		age[life]++
	}

	maxDays := 256
	day := 1

	//debugDay(0)

	for day <= maxDays {
		var tmpAge [9]int
		for i := 1; i < 9; i++ {
			tmpAge[i-1] = age[i]
		}
		tmpAge[8] += age[0]
		tmpAge[6] += age[0]
		age = tmpAge
		day++
	}

	total := 0
	for _, cnt := range age {
		total += cnt
	}

	log.Printf("Total Fish: %d\n", total)
}
