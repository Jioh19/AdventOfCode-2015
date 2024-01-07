package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type Rect struct {
	length int
	width  int
	height int
}

func main() {
	fileName := "input.txt"
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	s := time.Now()
	data := insertData(file)
	part1 := part1(data)
	fmt.Println("Part 1:", part1)
	fmt.Println("Time in nanoseconds:", time.Since(s).Nanoseconds())
	s = time.Now()
	part2 := part2(data)
	fmt.Println("Part 2:", part2)
	fmt.Println("Time in nanoseconds:", time.Since(s).Nanoseconds())
}

func insertData(file []byte) []Rect {
	rects := []Rect{}
	stringify := strings.TrimSpace(string(file))
	lines := strings.Split(stringify, "\n")
	for _, line := range lines {
		data := strings.Split(line, "x")
		len, _ := strconv.Atoi(data[0])
		wid, _ := strconv.Atoi(data[1])
		hei, _ := strconv.Atoi(data[2])
		rect := Rect{len, wid, hei}
		rects = append(rects, rect)
	}
	return rects
}

func part1(data []Rect) int {
	result := 0
	for _, rect := range data {
		area1 := rect.length * rect.width
		area2 := rect.height * rect.length
		area3 := rect.height * rect.width
		result += 2 * area1
		result += 2 * area2
		result += 2 * area3
		min := math.Min(float64(area1), float64(area2))
		min = math.Min(min, float64(area3))
		result += int(min)
	}
	return result
}

func part2(data []Rect) int {
	result := 0
	for _, rect := range data {
		len1, len2 := 0, 0
		if rect.length < rect.width {
			len1 = rect.length
			if rect.width < rect.height {
				len2 = rect.width
			} else {
				len2 = rect.height
			}
		} else {
			len1 = rect.width
			if rect.length < rect.height {
				len2 = rect.length
			} else {
				len2 = rect.height
			}
		}

		result += 2 * (len1 + len2)
		result += rect.width * rect.height * rect.length
	}
	return result
}
