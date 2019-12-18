package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func doRun(program []int, input []int) (output []int) {
	memory := make([]int, len(program))
	copy(memory, program)

	ip := 0
	for {
		instruction := memory[ip]
		opcode := instruction % 100

		getParameter := func(offset int) int {
			parameter := memory[ip+offset]
			mode := instruction / pow(10, offset+1) % 10
			switch mode {
			case 0: // position mode
				return memory[parameter]
			case 1: // immediate mode
				return parameter
			default:
				panic(fmt.Sprintf("fault: invalid parameter mode: ip=%d instruction=%d offset=%d mode=%d", ip, instruction, offset, mode))
			}
		}

		switch opcode {

		case 1: // ADD
			a, b, c := getParameter(1), getParameter(2), memory[ip+3]
			memory[c] = a + b
			ip += 4

		case 2: // MULTIPLY
			a, b, c := getParameter(1), getParameter(2), memory[ip+3]
			memory[c] = a * b
			ip += 4

		case 3: // INPUT
			a := memory[ip+1]
			memory[a] = input[0]
			input = input[1:]
			ip += 2

		case 4: // OUTPUT
			a := getParameter(1)
			output = append(output, a)
			ip += 2

		case 5: // JUMP IF TRUE
			a, b := getParameter(1), getParameter(2)
			if a != 0 {
				ip = b
			} else {
				ip += 3
			}

		case 6: // JUMP IF FALSE
			a, b := getParameter(1), getParameter(2)
			if a == 0 {
				ip = b
			} else {
				ip += 3
			}

		case 7: // LESS THAN
			a, b, c := getParameter(1), getParameter(2), memory[ip+3]
			if a < b {
				memory[c] = 1
			} else {
				memory[c] = 0
			}
			ip += 4

		case 8: // EQUAL
			a, b, c := getParameter(1), getParameter(2), memory[ip+3]
			if a == b {
				memory[c] = 1
			} else {
				memory[c] = 0
			}
			ip += 4

		case 99: // HALT
			return output

		default:
			panic(fmt.Sprintf("fault: invalid opcode: ip=%d instruction=%d opcode=%d", ip, instruction, opcode))
		}
	}
}

func main() {
	input := readFile("test.txt")

	var program []int
	for _, value := range strings.Split(input, ",") {
		program = append(program, toInt(value))
	}

	{
		fmt.Println("--- Part One ---")
		var phaseSettings [5][5]int
		for i := 0; i < 5; i++ {
			phase1 := i
			phaseSettings[i] = [5]int{phase1, (phase1 + 1) % 5, (phase1 + 2) % 5, (phase1 + 3) % 5, (phase1 + 4) % 5}
		}
		maxThruster := 0
		input := 0
		for i := range phaseSettings {
			for _, j := range phaseSettings[i] {
				output := doRun(program, []int{j, input})
				input = output[len(output)-1]
			}
			maxThruster = max(maxThruster, input)
			input = 0
		}
		fmt.Println(maxThruster)
	}

	//{
	//	fmt.Println("--- Part Two ---")
	//	output := doRun(program, []int{5})
	//	for i := 0; i < len(output)-1; i++ {
	//		if output[i] != 0 {
	//			panic(fmt.Sprintf("test failure: %v", output))
	//		}
	//	}
	//	fmt.Println(output[len(output)-1])
	//}
}

func pow(a, b int) int {
	p := 1
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

func toInt(s string) int {
	result, err := strconv.Atoi(s)
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
