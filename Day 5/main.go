package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fileName := "input.txt"
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	input := insertData(file)

	s := time.Now()
	part1 := part1(input)
	fmt.Println("Part 1:", part1)
	fmt.Println("Time in nanoseconds:", time.Since(s).Nanoseconds())
	s2 := time.Now()
	part2 := part2(input)
	fmt.Println("Part 2:", part2)
	fmt.Println("Time in milliseconds:", time.Since(s2).Milliseconds())
}

func insertData(file []byte) [][]byte {
	var input [][]byte
	//var word []byte
	lines := strings.Split(strings.TrimSpace(string(file)), "\n")
	for _, line := range lines {
		input = append(input, []byte(line))
	}
	return input
}

func checkVowel(line []byte) bool {
	count := 0
	for _, c := range line {
		switch c {
		case 'a', 'e', 'i', 'o', 'u':
			count++
		}
		if count >= 3 {
			return true
		}
	}
	return false
}

func checkTwin(line []byte) bool {
	for i := 0; i < len(line)-1; i++ {
		if line[i] == line[i+1] {
			return true
		}
	}
	return false
}

func checkContain(line []byte) bool {
	if strings.Contains(string(line), "ab") || strings.Contains(string(line), "cd") {
		return false
	}
	if strings.Contains(string(line), "pq") || strings.Contains(string(line), "xy") {
		return false
	}
	return true
}

func part1(lines [][]byte) int {
	result := 0
	for _, line := range lines {
		if checkContain(line) && checkTwin(line) && checkVowel(line) {
			result++
		}
	}
	return result
}

func checkSpace(line []byte) bool {
	for i := 0; i < len(line)-2; i++ {
		if line[i] == line[i+2] {
			return true
		}
	}
	return false
}

func checkPair(line []byte) bool {
	for i := 0; i < len(line)-3; i++ {
		pair := []byte{line[i], line[i+1]}
		if strings.Contains(string(line[i+2:]), string(pair)) {
			return true
		}
	}
	return false
}

func part2(lines [][]byte) int {
	result := 0
	for _, line := range lines {
		if checkPair(line) && checkSpace(line) {
			result++
		}
	}
	return result
}
