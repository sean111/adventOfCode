package main

import (
	"bufio"
	"fmt"
	"os"
)

type Rucksack struct {
	Compartment1 []rune
	Compartment2 []rune
}

func main() {
	data, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(data)

	var total int
	for scanner.Scan() {
		sack := ParseContainer(scanner.Text())
		dupes := FindDuplicates(sack)
		for _, dupe := range dupes {
			total += GetItemPriority(dupe)
		}
	}
	fmt.Printf("Total: %d\n", total)
}

func ParseContainer(items string) Rucksack {
	tmp := []rune(items)
	half := len(tmp) / 2
	c1Temp := tmp[0:half]
	c2Temp := tmp[half:(half * 2)]
	sack := Rucksack{Compartment1: c1Temp, Compartment2: c2Temp}
	return sack
}

func FindDuplicates(sack Rucksack) []rune {
	var duplicates []rune
	for _, c1Item := range sack.Compartment1 {
		if containsRune(sack.Compartment2, c1Item) && !containsRune(duplicates, c1Item) {
			duplicates = append(duplicates, c1Item)
		}
	}
	return duplicates
}

func GetItemPriority(item rune) int {
	tmp := int(item)
	var modifier int
	if tmp > 96 {
		modifier = 96
	} else {
		modifier = 38
	}
	return tmp - modifier
}

func containsRune(list []rune, target rune) bool {
	for _, item := range list {
		if item == target {
			return true
		}
	}
	return false
}
