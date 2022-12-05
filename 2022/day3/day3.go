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

func (sack Rucksack) Combined() []rune {
	return append(sack.Compartment1, sack.Compartment2...)
}

func (sack Rucksack) Contains(item rune) bool {
	//Check compartment 1
	sackItems := sack.Combined()
	for _, sackItem := range sackItems {
		if sackItem == item {
			return true
		}
	}
	return false
}

func main() {
	data, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(data)

	var total int
	var p2Total int
	groupCount := 0
	group := make([]Rucksack, 0)
	for scanner.Scan() {
		sack := ParseContainer(scanner.Text())
		dupes := FindDuplicates(sack)
		for _, dupe := range dupes {
			total += GetItemPriority(dupe)
		}
		groupCount++
		if groupCount == 3 {
			// Do group check logic
			commonItem := CheckGroupForCommonItem(append(group, sack))
			if commonItem != rune(0) {
				p2Total += GetItemPriority(commonItem)
			}
			groupCount = 0
			group = make([]Rucksack, 0)
		} else {
			group = append(group, sack)
		}
	}
	// Check the end of the group array
	fmt.Printf("P1 Total: %d\n", total)
	fmt.Printf("P2 Total: %d\n", p2Total)
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

func CheckGroupForCommonItem(sacks []Rucksack) rune {
	firstSack := sacks[0].Combined()
	for _, item := range firstSack {
		if sacks[1].Contains(item) && sacks[2].Contains(item) {
			return item
		}
	}
	return rune(0)
}

func containsRune(list []rune, target rune) bool {
	for _, item := range list {
		if item == target {
			return true
		}
	}
	return false
}
