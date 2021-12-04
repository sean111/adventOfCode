package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

type PosInfo struct {
	totalZeros int
	totalOnes  int
}

func main() {
	var positionArray []PosInfo
	var entryArray []string
	var oxyArray []string
	var cdoxArray []string
	var strLen int
	gammaRate := ""
	epsilonRate := ""
	data, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer data.Close()
	buffer := bufio.NewScanner(data)
	for buffer.Scan() {
		entry := buffer.Text()
		entryArray = append(entryArray, entry)
		strLen = len(entry)

		// log.Printf("Entry: %s | Length: %d\n", entry, strLen)

		for x := 0; x < strLen; x++ {
			if x == len(positionArray) {
				positionArray = append(positionArray, PosInfo{totalOnes: 0, totalZeros: 0})
			}
			if posValue(entry, x) == "1" {
				positionArray[x].totalOnes++
			} else {
				positionArray[x].totalZeros++
			}
		}
	}

	for x := 0; x < strLen; x++ {
		if positionArray[x].totalOnes > positionArray[x].totalZeros {
			gammaRate += "1"
			epsilonRate += "0"
		} else {
			gammaRate += "0"
			epsilonRate += "1"
		}
	}

	gammaDec, _ := strconv.ParseInt(gammaRate, 2, 64)
	epsilonDec, _ := strconv.ParseInt(epsilonRate, 2, 64)

	product := gammaDec * epsilonDec
	log.Printf("gammaRate: %s (%d)| epsilonRate: %s (%d)\n", gammaRate, gammaDec, epsilonRate, epsilonDec)
	log.Printf("Part1 Answer: %d", product)

	//Part 2
	oxyArray = entryArray
	cdoxArray = entryArray

	for x := 0; x < strLen; x++ {
		oxyZeros, oxyOnes := posCount(oxyArray, x)
		var oxyTarget string
		if oxyZeros == oxyOnes {
			oxyTarget = "1"
		} else if oxyZeros > oxyOnes {
			oxyTarget = "0"
		} else {
			oxyTarget = "1"
		}
		oxyArray = filterByPosValue(oxyArray, x, oxyTarget)

		cdoxZeros, cdoxOnes := posCount(cdoxArray, x)
		var cdoxTarget string

		if cdoxZeros == cdoxOnes {
			cdoxTarget = "0"
		} else if cdoxZeros > cdoxOnes {
			cdoxTarget = "1"
		} else {
			cdoxTarget = "0"
		}
		cdoxArray = filterByPosValue(cdoxArray, x, cdoxTarget)
	}

	//log.Printf("%#v", oxyArray)
	//log.Printf("%#v", cdoxArray)

	oxyFinalValue, _ := strconv.ParseInt(oxyArray[0], 2, 64)
	cdoxFinalValue, _ := strconv.ParseInt(cdoxArray[0], 2, 64)

	part2Final := oxyFinalValue * cdoxFinalValue

	log.Printf("Part2 Answer: %d\n", part2Final)
}

func posValue(target string, position int) string {
	return string([]rune(target)[position])
}

func posCount(values []string, position int) (int, int) {
	var totalZeros = 0
	var totalOnes = 0
	for _, value := range values {
		if posValue(value, position) == "0" {
			totalZeros++
		} else {
			totalOnes++
		}
	}
	return totalZeros, totalOnes
}

func filterByPosValue(values []string, position int, targetValue string) []string {
	if len(values) == 1 {
		return values
	}
	var tmpArray []string
	for _, value := range values {
		if posValue(value, position) == targetValue {
			tmpArray = append(tmpArray, value)
		}
	}
	return tmpArray
}
