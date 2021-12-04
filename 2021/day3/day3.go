package main

import (
	"os"
	"log"
	"bufio"
	"strconv"
)

type PosInfo struct {
	totalZeros int
	totalOnes int
}

func main() {
	var positionArray []PosInfo
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
		strLen = len(entry)

		// log.Printf("Entry: %s | Length: %d\n", entry, strLen)

		for x := 0; x < strLen; x++ {
			if x == len(positionArray) {
				positionArray = append(positionArray, PosInfo{totalOnes: 0, totalZeros: 0})
			}
			if (string([]rune(entry)[x]) == "1") {
				positionArray[x].totalOnes++
			} else {
				positionArray[x].totalZeros++
			}
		}
	}

	for x := 0; x < strLen; x++ {
		if (positionArray[x].totalOnes > positionArray[x].totalZeros) {
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

	// log.Printf("%#v", positionArray)
	log.Printf("gammaRate: %s (%d)| epsilonRate: %s (%d)\n", gammaRate, gammaDec,  epsilonRate, epsilonDec)
	log.Printf("Part1 Answer: %d", product)

}