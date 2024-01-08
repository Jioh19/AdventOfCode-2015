package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"time"
)

func main() {
	data := []byte("abcdefg")
	len := len(data)
	hash := fmt.Sprintf("%x", md5.Sum(data))
	s := time.Now()
	for i := 0; hash[:5] != "00000"; i++ {
		num := []byte(strconv.Itoa(i))
		data = append(data[:len], num...)
		hash = fmt.Sprintf("%x", md5.Sum(data))
	}

	fmt.Println("Part 1:", string(data))
	fmt.Println("Time in nanoseconds:", time.Since(s).Nanoseconds())

	s = time.Now()
	for i := 0; hash[:6] != "000000"; i++ {
		num := []byte(strconv.Itoa(i))
		data = append(data[:len], num...)
		hash = fmt.Sprintf("%x", md5.Sum(data))
	}

	fmt.Println("Part 2:", string(data))
	fmt.Println("Time in nanoseconds:", time.Since(s).Nanoseconds())

}
