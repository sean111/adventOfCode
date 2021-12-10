package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Best struct {
	position int
	cost     int
}

func main() {
	var crabs []int
	var posCost int
	var maxPos = 0
	best := Best{position: 0, cost: -1}
	data, err := os.ReadFile("data.txt")
	if err != nil {
		log.Fatalln(err)
		return
	}
	initialData := strings.Split(strings.TrimSuffix(string(data), "\n"), ",")
	log.Printf("%#v\n", initialData)

	for _, val := range initialData {
		tmp, _ := strconv.Atoi(val)
		crabs = append(crabs, tmp)
		if tmp > maxPos {
			maxPos = tmp
		}
	}

	log.Printf("%#v", crabs)

	for pos := 0; pos < maxPos; pos++ {
		posMarker := pos + 1
		posCost = 0
		log.Printf("Checking %d", posMarker)
		for _, val := range crabs {
			moves := math.Abs(float64(val) - float64(pos))
			posCost += fuelCost(int(moves))
		}

		log.Printf("Position: %d | Cost: %d", pos, posCost)

		if posCost < best.cost || best.cost == -1 {
			log.Printf("New best pos: %d | cost: %d", posMarker, posCost)
			best = Best{position: posMarker, cost: posCost}
		}
	}

	fmt.Printf("Best Move: %d (%d)", best.position, best.cost)
}

func fuelCost(moves int) int {
	var cost int
	for i := 1; i <= moves; i++ {
		cost += i
	}
	return cost
}
