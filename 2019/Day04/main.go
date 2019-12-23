package main

import (
	"fmt"
	"strconv"
	"strings"
)

func passwordCount(min int, max int) int {
	count := 0
	for input := min; input < max; input++ {
		chars := strings.Split(strconv.Itoa(input), "")
		isIncreasing := true
		doubles := make(map[string]int)
		for i, v := range chars {
			if i < len(chars)-1 {
				val1 := v
				val2 := chars[i+1]
				isIncreasing = isIncreasing && val2 >= val1
				if val1 == val2 {
					_, exist := doubles[val1]

					if exist {
						doubles[val1]++
					} else {
						doubles[val1] = 2
					}
				}
			}
		}
		hasDouble := false
		for _, v := range doubles {
			if v == 2 {
				hasDouble = true
				break
			}
		}
		if isIncreasing && hasDouble {
			count++
		}
	}
	return count
}

func toInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func main() {
	min := 236491
	max := 713787
	count := passwordCount(min, max)
	fmt.Println("Number of passwords: ", count)
}
