package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

//type File struct {
//	name string
//	size int
//}

type Directory struct {
	name   string
	parent *Directory
	dirs   []Directory
	files  map[string]int
}

var DirSizes []int

func main() {
	data, err := os.Open("data.txt")
	handleError(err)

	scanner := bufio.NewScanner(data)
	fileSystem := Directory{name: "root"}
	currentDir := &fileSystem
	for scanner.Scan() {
		input := strings.Fields(scanner.Text())
		log.Printf("Input: %v\n", input)
		switch input[0] {
		case "$":
			// Command
			if len(input) > 2 {
				if input[2] == ".." {
					//Go up a dir
					currentDir = currentDir.parent
				} else {
					currentDir.dirs = append(currentDir.dirs, Directory{name: input[2], parent: currentDir, files: make(map[string]int, 0)})
					currentDir = &currentDir.dirs[len(currentDir.dirs)-1]
				}
			}
		case "dir":
			continue
		default:
			size, _ := strconv.Atoi(input[0])
			currentDir.files[input[1]] = size
		}
	}

	printDir(fileSystem)
	totalSize := getDirSize(fileSystem)

	log.Printf("Before: %v\n", DirSizes)
	sort.Ints(DirSizes)
	log.Printf("After: %v\n", DirSizes)

	p1Answer := 0
	p2Answer := 0
	for _, size := range DirSizes {
		if size < 100000 {
			p1Answer += size
		}
		if (70000000-totalSize)+size > 30000000 {
			p2Answer = size
			break
		}
	}
	fmt.Printf("Total Size: %d\n", totalSize)
	fmt.Printf("P1 Answer: %d\n", p1Answer)
	fmt.Printf("P2 Answer: %d\n", p2Answer)
}

func printDir(dir Directory) {
	if dir.parent != nil {
		fmt.Printf("Parent: %s || ", dir.parent.name)
	}
	fmt.Printf("Dir: %s\n", dir.name)
	for file, size := range dir.files {
		log.Printf("%s - %d", file, size)
	}
	for _, subDir := range dir.dirs {
		printDir(subDir)
	}
}

func getDirSize(dir Directory) int {
	totalSize := 0
	for _, size := range dir.files {
		totalSize += size
	}
	for _, subDir := range dir.dirs {
		totalSize += getDirSize(subDir)
	}
	DirSizes = append(DirSizes, totalSize)
	log.Printf("Dir: %s - Size: %d\n", dir.name, totalSize)
	return totalSize
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
