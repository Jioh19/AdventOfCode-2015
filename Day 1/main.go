package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fileName := "input.txt"
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	// for i, line := range grid {
	// 	fmt.Println(i, string(line))
	// }
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
	for _, c := range file {
		if c == 40 {
			result++
		}
		if c == 41 {
			result--
		}
	}
	return result
}

func part2(file []byte) int {
	result := 0
	for i, c := range file {
		if c == 40 {
			result++
		}
		if c == 41 {
			result--
		}
		if result == -1 {
			return i + 1
		}
	}
	return result
}

// func insertData(file []byte) []byte {
// 	var input [][]byte
// 	var word []byte
// 	for _, letter := range file {
// 		if letter != 13 {

// 			if letter == '\n' {
// 				input = append(input, word)
// 				word = []byte{}
// 			} else {
// 				word = append(word, letter)
// 			}
// 		}
// 	}
// 	return input
// }
