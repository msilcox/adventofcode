package main

import (
	"fmt"
)

func runProgram(input []int, batch []int) {
	opcode := batch[0]
	pos := batch[3]
	output := 0
	if opcode == 1 {
		output = input[batch[1]] + input[batch[2]]
	} else {
		output = input[batch[1]] * input[batch[2]]
	}
	input[pos] = output
}

func doRun(noun int, verb int) int {
	input := []int{1, noun, verb, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 10, 1, 19, 1, 19, 9, 23, 1, 23, 6, 27, 2, 27, 13, 31, 1, 10, 31, 35, 1, 10, 35, 39, 2, 39, 6, 43, 1, 43, 5, 47, 2, 10, 47, 51, 1, 5, 51, 55, 1, 55, 13, 59, 1, 59, 9, 63, 2, 9, 63, 67, 1, 6, 67, 71, 1, 71, 13, 75, 1, 75, 10, 79, 1, 5, 79, 83, 1, 10, 83, 87, 1, 5, 87, 91, 1, 91, 9, 95, 2, 13, 95, 99, 1, 5, 99, 103, 2, 103, 9, 107, 1, 5, 107, 111, 2, 111, 9, 115, 1, 115, 6, 119, 2, 13, 119, 123, 1, 123, 5, 127, 1, 127, 9, 131, 1, 131, 10, 135, 1, 13, 135, 139, 2, 9, 139, 143, 1, 5, 143, 147, 1, 13, 147, 151, 1, 151, 2, 155, 1, 10, 155, 0, 99, 2, 14, 0, 0}
	toBatch := input
	batchSize := 4
	var batches [][]int

	for batchSize < len(toBatch) {
		toBatch, batches = toBatch[batchSize:], append(batches, toBatch[0:batchSize:batchSize])
	}
	batches = append(batches, toBatch)
	for _, v := range batches {
		if v[0] == 99 {
			break
		}
		runProgram(input, v)
	}
	return input[0]
}

func main() {
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			if doRun(i, j) == 19690720 {
				ans := 100*i + j
				fmt.Println(ans)
			}
		}
	}
}
