package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Fish struct {
	age int
}

var school []Fish

func main() {
	data, err := os.ReadFile("data.txt")
	if err != nil {
		log.Fatalln(err)
		return
	}
	initialState := strings.Split(strings.TrimSuffix(string(data), "\n"), ",")

	for _, val := range initialState {
		life, _ := strconv.Atoi(val)
		school = append(school, Fish{age: life})
	}

	maxDays := 256
	day := 1

	//debugDay(0)

	for day <= maxDays {
		for i, _ := range school {
			fish := school[i]
			if fish.age == 0 {
				fish.age = 6
				school = append(school, Fish{age: 8})
			} else {
				fish.age--
			}
			school[i] = fish
		}
		//debugDay(day)
		day++
	}
	log.Printf("Total Fish: %d", len(school))
}

func debugDay(day int) {
	fmt.Printf("Day %d:     ", day)
	for _, fish := range school {
		fmt.Printf("%d, ", fish.age)
	}
	fmt.Println("")
}
