package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	count := 0
	file, err := os.Open("/Users/marksilcox/Documents/Work/GO/adventofcode/2021/Day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	prev := -1
	for sc.Scan() {
		i, err := strconv.Atoi(sc.Text())
		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		}
		if prev != -1 {
			if prev < i {
				count = count + 1
			}
		}
		prev = i
	}

	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(fmt.Sprintf("Count = %d", count))
}
