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
	part2 := part2(file)
	fmt.Println("Part 2:", part2)
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
	result = len(myMap)
	return result
}

func part2(file []byte) int {
	result := 0
	var (
		i1 int
		i2 int
		j1 int
		j2 int
	)
	coord := Coord{i1, j1}
	myMap := make(map[Coord]int)
	myMap[coord] += 1
	for r, c := range file {
		if r%2 == 0 {
			i1, j1 = getCoord(i1, j1, c)
			coord = Coord{i1, j1}
		} else {
			i2, j2 = getCoord(i2, j2, c)
			coord = Coord{i2, j2}
		}
		myMap[coord] += 1
	}
	result = len(myMap)
	return result
}

func getCoord(i, j int, c byte) (int, int) {
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
	return i, j
}
