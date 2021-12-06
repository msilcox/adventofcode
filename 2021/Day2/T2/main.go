package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("/Users/marksilcox/Documents/Work/GO/adventofcode/2021/Day2/input.txt")
	defer file.Close()

	sc := bufio.NewScanner(file)
	horiz := 0
	depth := 0
	aim := 0

	for sc.Scan() {
		cmd := strings.Fields(sc.Text())
		val, _ := strconv.Atoi(cmd[1])
		if cmd[0] == "down" {
			aim += val
		}
		if cmd[0] == "up" {
			aim -= val
		}
		if cmd[0] == "forward" {
			horiz += val
			depth += aim * val
		}
	}
	fmt.Println(fmt.Sprintf("Answer = %d", horiz*depth))
}
