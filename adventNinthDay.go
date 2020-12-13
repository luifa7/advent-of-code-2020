package main

import (
	"fmt"
	"sort"
)

var pathToInputNinthDay = pathToInputs + "09"
var codes = convertSSIntoSI(readFileLines(pathToInputNinthDay))

var preamble = 25

var wrongXMAS = 90433990

// var wrongXMAS = 127

func resolveXmasWrongCode() {
	sn := findSequenceForWrongXmas()
	sort.Ints(sn)
	fmt.Println(sn)
	fmt.Println(sn[0] + sn[len(sn)-1])
}

func findSequenceForWrongXmas() []int {
	startOn := 0
	sum := 0
	for startOn < len(codes) {
		sc := codes[startOn:]
		for i, n := range sc {
			sum += n
			if sum == wrongXMAS {
				sn := sc[:i+1]
				return sn
			} else if sum > wrongXMAS {
				startOn++
				sum = 0
				break
			}
		}
	}
	return []int{}
}

func findWrongXMAS() {

	for i := preamble; i < len(codes); i++ {
		n := codes[i]
		sn := codes[i-preamble : i]
		if !isNumSumOfTwoOfThem(n, sn) {
			fmt.Println("Not XMAS:", n)
			break
		}
	}
}

func isNumSumOfTwoOfThem(n int, sn []int) bool {
	for _, n1 := range sn {
		for _, n2 := range sn {
			if n1 != n2 && (n1+n2) == n {
				return true
			}
		}
	}
	return false
}
