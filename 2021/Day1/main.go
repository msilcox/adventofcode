package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("/Users/marksilcox/Documents/Work/GO/adventofcode/2021/Day1/input.txt")
	defer file.Close()

	sc := bufio.NewScanner(file)
	prev := -1
	counter := 0
	var buffer []int
	for sc.Scan() {
		i, _ := strconv.Atoi(sc.Text())
		buffer = append(buffer, i)
		if len(buffer) < 3 {
			continue
		}
		newSum := sum(buffer)
		if prev != -1 {
			if newSum > prev {
				counter++
			}
		}
		prev = newSum
		buffer = buffer[1:]
	}

	fmt.Println(fmt.Sprintf("Count = %d", counter))
}

func sum(buffer []int) int {
	result := 0
	for _, v := range buffer {
		result += v
	}
	return result
}
