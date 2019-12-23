package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	chunks := readInput()

	{
		fmt.Println("--- Part One ---")
		checksum(chunks)
	}

	{
		fmt.Println("--- Part Two ---")
		readImage(chunks)
	}
}

func readImage(chunks []string) {
	var image string
	for _, chunk := range chunks {
		image = buildImage(image, chunk)
	}
	finalImage := chunkString(image, 25)
	for _, line := range finalImage {
		fmt.Println(line)
	}
}

func buildImage(s1 string, s2 string) string {
	if len(s1) > 0 {
		r1 := []rune(s1)
		r2 := []rune(s2)
		out := bytes.Buffer{}
		for i := 0; i < len(r1); i++ {
			if r1[i] == []rune("2")[0] {
				out.WriteRune(getRune(r2[i]))
			} else {
				out.WriteRune(getRune(r1[i]))
			}
		}
		return out.String()
	} else {
		return s2
	}
}

func getRune(r rune) rune {
	vals := []rune("012")
	if r == vals[0] {
		return []rune(" ")[0]
	} else if r == vals[1] {
		return []rune("X")[0]
	} else {
		return r
	}
}

func readInput() []string {
	input := readFile("input.txt")
	var chunks []string
	chunkSize := 25 * 6
	chunks = chunkString(input, chunkSize)
	return chunks
}

func chunkString(input string, chunkSize int) []string {
	var chunks []string
	runes := []rune(input)
	for i := 0; i < len(runes); i += chunkSize {
		nn := i + chunkSize
		if nn > len(runes) {
			nn = len(runes)
		}
		chunks = append(chunks, string(runes[i:nn]))
	}
	return chunks
}

func checksum(chunks []string) {
	var foundChunk string
	minZeros := math.MaxInt32
	for _, chunk := range chunks {
		zero := regexp.MustCompile("0")

		zeroCount := len(zero.FindAllString(chunk, -1))
		if zeroCount < minZeros {
			minZeros = zeroCount
			foundChunk = chunk
		}
	}
	one := regexp.MustCompile("1")
	oneCount := len(one.FindAllString(foundChunk, -1))
	two := regexp.MustCompile("2")
	twoCount := len(two.FindAllString(foundChunk, -1))
	fmt.Println(oneCount * twoCount)
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

func min(x, y int) int {

	return x
}
