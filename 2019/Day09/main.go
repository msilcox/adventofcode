package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func doRun(program []int64, input []int64) (output []int64) {
	memory := make([]int64, len(program))
	copy(memory, program)

	getMemoryPointer := func(index int64) *int64 {
		// Grow memory, if index is out of range.
		for int64(len(memory)) <= index {
			memory = append(memory, 0)
		}
		return &memory[index]
	}

	var ip, relativeBase int64
	for {
		instruction := memory[ip]
		opcode := instruction % 100

		getParameter := func(offset int64) *int64 {
			parameter := memory[ip+offset]
			mode := instruction / pow(10, offset+1) % 10
			switch mode {
			case 0: // position mode
				return getMemoryPointer(parameter)
			case 1: // immediate mode
				return &parameter
			case 2: // relative mode
				return getMemoryPointer(relativeBase + parameter)
			default:
				panic(fmt.Sprintf("fault: invalid parameter mode: ip=%d instruction=%d offset=%d mode=%d", ip, instruction, offset, mode))
			}
		}

		switch opcode {

		case 1: // ADD
			a, b, c := getParameter(1), getParameter(2), getParameter(3)
			*c = *a + *b
			ip += 4

		case 2: // MULTIPLY
			a, b, c := getParameter(1), getParameter(2), getParameter(3)
			*c = *a * *b
			ip += 4

		case 3: // INPUT
			a := getParameter(1)
			*a = input[0]
			input = input[1:]
			ip += 2

		case 4: // OUTPUT
			a := getParameter(1)
			output = append(output, *a)
			ip += 2

		case 5: // JUMP IF TRUE
			a, b := getParameter(1), getParameter(2)
			if *a != 0 {
				ip = *b
			} else {
				ip += 3
			}

		case 6: // JUMP IF FALSE
			a, b := getParameter(1), getParameter(2)
			if *a == 0 {
				ip = *b
			} else {
				ip += 3
			}

		case 7: // LESS THAN
			a, b, c := getParameter(1), getParameter(2), getParameter(3)
			if *a < *b {
				*c = 1
			} else {
				*c = 0
			}
			ip += 4

		case 8: // EQUAL
			a, b, c := getParameter(1), getParameter(2), getParameter(3)
			if *a == *b {
				*c = 1
			} else {
				*c = 0
			}
			ip += 4

		case 9: // RELATIVE BASE OFFSET
			a := getParameter(1)
			relativeBase += *a
			ip += 2

		case 99: // HALT
			return output

		default:
			panic(fmt.Sprintf("fault: invalid opcode: ip=%d instruction=%d opcode=%d", ip, instruction, opcode))
		}
	}
}

func main() {
	input := readFile("input.txt")

	var program []int64
	for _, value := range strings.Split(input, ",") {
		program = append(program, toInt(value))
	}

	{
		fmt.Println("--- Part One ---")

		output := doRun(program, []int64{1})
		for i := 0; i < len(output)-1; i++ {
			if output[i] != 0 {
				panic(fmt.Sprintf("test failure: %v", output))
			}
		}
		fmt.Println(output[len(output)-1])
	}

	{
		fmt.Println("--- Part Two ---")

		output := doRun(program, []int64{2})
		for i := 0; i < len(output)-1; i++ {
			if output[i] != 0 {
				panic(fmt.Sprintf("test failure: %v", output))
			}
		}
		fmt.Println(output[len(output)-1])
	}
}

func pow(a, b int64) int64 {
	p := int64(1)
	for b > 0 {
		if b&1 != 0 {
			p *= a
		}
		b >>= 1
		a *= a
	}
	return p
}

func readFile(filename string) string {
	bytes, err := ioutil.ReadFile(filename)
	check(err)
	return strings.TrimSpace(string(bytes))
}

func toInt(s string) int64 {
	result, err := strconv.ParseInt(s, 10, 64)
	check(err)
	return result
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func max(x, y int) int {
	if y > x {
		return y
	}
	return x
}
