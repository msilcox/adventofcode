package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func emulate(program []int, input <-chan int, output chan<- int, halt chan<- bool) {
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
			memory[a] = <-input
			ip += 2

		case 4: // OUTPUT
			a := getParameter(1)
			output <- a
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
			halt <- true
			return

		default:
			panic(fmt.Sprintf("fault: invalid opcode: ip=%d instruction=%d opcode=%d", ip, instruction, opcode))
		}
	}
}

func main() {
	input := readFile("input.txt")

	var program []int
	for _, value := range strings.Split(input, ",") {
		program = append(program, toInt(value))
	}

	{
		fmt.Println("--- Part One ---")

		phaseSettings := getPermutations([]int{0,1,2,3,4})
		maxThruster := 0
		for _, p := range phaseSettings {
             e := emulateAmplifiers(program, p)

			maxThruster = max(maxThruster, e)
		}
		fmt.Println(maxThruster)
	}

	{
		fmt.Println("--- Part Two ---")

	    phaseSettings := getPermutations([]int{5,6,7,8,9})
    	maxThruster := 0
    	for _, p := range phaseSettings {
                e := emulateAmplifiers(program, p)

    			maxThruster = max(maxThruster, e)
    		}
    		fmt.Println(maxThruster)
	}
}

func emulateAmplifiers(program []int, phaseSettings []int) int {
	// Set up the channels connecting the amplifiers.
	ea := make(chan int, 1) // must be buffered to receive final result
	ab := make(chan int)
	bc := make(chan int)
	cd := make(chan int)
	de := make(chan int)

	// This channel will receive a value each time an amplifier halts.
	halt := make(chan bool)

	// Start amplifiers in parallel.
	go emulate(program, ea, ab, halt)
	go emulate(program, ab, bc, halt)
	go emulate(program, bc, cd, halt)
	go emulate(program, cd, de, halt)
	go emulate(program, de, ea, halt)

	// Provide phase settings.
	ea <- phaseSettings[0]
	ab <- phaseSettings[1]
	bc <- phaseSettings[2]
	cd <- phaseSettings[3]
	de <- phaseSettings[4]

	// Send initial input signal.
	ea <- 0

	// Wait for all amplifiers to halt.
	for i := 0; i < 5; i++ {
		<-halt
	}

	// Read the final result.
	return <-ea
}

func getPermutations(elements []int) [][]int {
	permutations := [][]int{}
	if len(elements) == 1 {
		permutations = [][]int{elements}
		return permutations
	}
	for i := range elements {
		el := make([]int, len(elements))
		copy(el, elements)

		for _, perm := range getPermutations(append(el[0:i], el[i+1:]...)) {
			permutations = append(permutations, append([]int{elements[i]}, perm...))
		}
	}
	return permutations
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
