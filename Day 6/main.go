package main

import (
	"fmt"
	"os"
	"strconv"
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
	lines := insertData(file)
	s := time.Now()
	part1 := part1(lines)
	fmt.Println("Part 1:", part1)
	fmt.Println("Time in nanoseconds:", time.Since(s).Nanoseconds())

	s = time.Now()
	part2 := part2(lines)
	fmt.Println("Part 2:", part2)
	fmt.Println("Time in nanoseconds:", time.Since(s).Nanoseconds())

}

func insertData(file []byte) []string {
	lines := (strings.Split(strings.TrimSpace(string(file)), "\n"))
	return lines
}

func part1(lines []string) int {
	result := 0
	lights := [1000][1000]bool{}
	for _, line := range lines {
		if strings.Contains(line, "turn on") {
			turnOn(&lights, strings.Split(line, "turn on ")[1])
		}
		if strings.Contains(line, "turn off") {
			turnOff(&lights, strings.Split(line, "turn off ")[1])
		}
		if strings.Contains(line, "toggle") {
			toggle(&lights, strings.Split(line, "toggle ")[1])
		}
	}
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if lights[i][j] {
				result++
			}
		}
	}
	return result
}

func turnOn(lights *[1000][1000]bool, line string) {
	inputs := strings.Split(line, " through ")
	start := strings.Split(inputs[0], ",")
	end := strings.Split(inputs[1], ",")
	start0, _ := strconv.Atoi(start[0])
	start1, _ := strconv.Atoi(start[1])
	end0, _ := strconv.Atoi(end[0])
	end1, _ := strconv.Atoi(end[1])
	for i := start0; i <= end0; i++ {
		for j := start1; j <= end1; j++ {
			lights[i][j] = true
		}
	}
}

func turnOff(lights *[1000][1000]bool, line string) {
	inputs := strings.Split(line, " through ")
	start := strings.Split(inputs[0], ",")
	end := strings.Split(inputs[1], ",")
	start0, _ := strconv.Atoi(start[0])
	start1, _ := strconv.Atoi(start[1])
	end0, _ := strconv.Atoi(end[0])
	end1, _ := strconv.Atoi(end[1])
	for i := start0; i <= end0; i++ {
		for j := start1; j <= end1; j++ {
			lights[i][j] = false
		}
	}
}

func toggle(lights *[1000][1000]bool, line string) {
	inputs := strings.Split(line, " through ")
	start := strings.Split(inputs[0], ",")
	end := strings.Split(inputs[1], ",")
	start0, _ := strconv.Atoi(start[0])
	start1, _ := strconv.Atoi(start[1])
	end0, _ := strconv.Atoi(end[0])
	end1, _ := strconv.Atoi(end[1])
	for i := start0; i <= end0; i++ {
		for j := start1; j <= end1; j++ {
			if lights[i][j] {
				lights[i][j] = false
			} else {
				lights[i][j] = true
			}
		}
	}
}

func part2(lines []string) int {
	result := 0
	lights := [1000][1000]int{}
	change := 0
	for _, line := range lines {
		if strings.Contains(line, "turn on") {
			change = 1
			changeLight(&lights, strings.Split(line, "turn on ")[1], change)
		}
		if strings.Contains(line, "turn off") {
			change = -1
			changeLight(&lights, strings.Split(line, "turn off ")[1], change)
		}
		if strings.Contains(line, "toggle") {
			change = 2
			changeLight(&lights, strings.Split(line, "toggle ")[1], change)
		}
	}
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {

			result += lights[i][j]

		}
	}
	return result
}

func changeLight(lights *[1000][1000]int, line string, change int) {
	inputs := strings.Split(line, " through ")
	start := strings.Split(inputs[0], ",")
	end := strings.Split(inputs[1], ",")
	start0, _ := strconv.Atoi(start[0])
	start1, _ := strconv.Atoi(start[1])
	end0, _ := strconv.Atoi(end[0])
	end1, _ := strconv.Atoi(end[1])
	for i := start0; i <= end0; i++ {
		for j := start1; j <= end1; j++ {
			lights[i][j] += change
			if lights[i][j] < 0 {
				lights[i][j] = 0
			}
		}
	}
}
