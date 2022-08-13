package main

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"os"
	"strconv"
)

func main() {
	secret := os.Args[1]
	number := 1
	for {
		test := []byte(secret + strconv.Itoa(number))
		hash := md5.Sum(test)
		output := hex.EncodeToString(hash[:])
		log.Printf("Testing %d (%s) (%s) (%s)\n", number, test, output, output[0:6])
		if output[0:6] == "000000" {
			break
		}
		number++
	}
	log.Printf("Answer: %d", number)
}
