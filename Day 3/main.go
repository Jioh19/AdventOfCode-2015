package main

import (
	"fmt"
	"os"
	"time"
)

type Coord struct {
	i int
	j int
}

func main() {
	fileName := "input.txt"
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	s := time.Now()
	part1 := part1(file)
	fmt.Println("Part 1:", part1)
	fmt.Println("Time in nanoseconds:", time.Since(s).Nanoseconds())
	s = time.Now()

	fmt.Println("Time in nanoseconds:", time.Since(s).Nanoseconds())
}

func part1(file []byte) int {
	result := 0
	i, j := 0, 0
	coord := Coord{i, j}
	myMap := make(map[Coord]int)
	myMap[coord] += 1
	for _, c := range file {
		switch c {
		case '^':
			i--
		case '>':
			j++
		case '<':
			j--
		case 'v':
			i++
		}
		coord = Coord{i, j}
		myMap[coord] += 1
	}
	fmt.Println(myMap)
	result = len(myMap)
	return result
}
