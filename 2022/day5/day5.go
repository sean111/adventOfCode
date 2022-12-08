package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Stack struct {
	boxes []string
}

type Cargo struct {
	columns []Stack
}

func (c Cargo) MovePos(count int, from int, to int) error {
	from--
	to--
	//log.Printf("[Start] From: %v || To: %v\n", c.columns[from], c.columns[to])
	length := len(c.columns[from].boxes)
	//log.Printf("Len: %d || Count: %d\n", length, count)
	if length < count {
		return fmt.Errorf("count is longer then the length [%d, %d]", count, length)
	}
	items := c.columns[from].boxes[length-count : length]
	for i := len(items) - 1; i >= 0; i-- {
		c.columns[to].boxes = append(c.columns[to].boxes, items[i])
	}
	c.columns[from].boxes = c.columns[from].boxes[0 : length-count]
	//log.Printf("Items: %v\n", items)
	//log.Printf("[End] From: %v || To: %v\n", c.columns[from], c.columns[to])
	return nil
}

func (c Cargo) GetTopCrates() string {
	var output string
	for _, column := range c.columns {
		output = fmt.Sprintf("%s%s", output, column.boxes[len(column.boxes)-1])
	}
	return output
}

func main() {

	cargo := Cargo{
		columns: []Stack{
			{
				boxes: []string{"Z", "J", "G"},
			},
			{
				boxes: []string{"Q", "L", "R", "P", "W", "F", "V", "C"},
			},
			{
				boxes: []string{"F", "P", "M", "C", "L", "G", "R"},
			},
			{
				boxes: []string{"L", "F", "B", "W", "P", "H", "M"},
			},
			{
				boxes: []string{"G", "C", "F", "S", "V", "Q"},
			},
			{
				boxes: []string{"W", "H", "J", "Z", "M", "Q", "T", "L"},
			},
			{
				boxes: []string{"H", "F", "S", "B", "V"},
			},
			{
				boxes: []string{"F", "J", "Z", "S"},
			},
			{
				boxes: []string{"M", "C", "D", "P", "F", "H", "B", "T"},
			},
		},
	}
	data, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(data)

	r := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
	for scanner.Scan() {
		commands := r.FindAllStringSubmatch(scanner.Text(), -1)

		count, _ := strconv.Atoi(commands[0][1])
		from, _ := strconv.Atoi(commands[0][2])
		to, _ := strconv.Atoi(commands[0][3])

		cargo.MovePos(count, from, to)

		//fmt.Printf("Command: %s || Parsed => Count: %d, From: %d, To: %d\n", scanner.Text(), count, from, to)
	}
	fmt.Printf("Top Containers: %s\n", cargo.GetTopCrates())
}
