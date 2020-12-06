package main

import (
	"fmt"
	"strconv"
	"strings"
)

var pathToInputSecondDay = pwd + "/inputs/input_02"

func parsePassLinesToStruct() []password {
	resutl := readFileLines(pathToInputSecondDay)
	var toReturn []password
	for _, r := range resutl {
		s := strings.Fields(r)
		rangeForLetter := strings.Split(s[0], "-")
		min, _ := strconv.Atoi(rangeForLetter[0])
		max, _ := strconv.Atoi(rangeForLetter[1])
		newPass := password{min, max, string(s[1][0]), s[2]}
		toReturn = append(toReturn, newPass)
	}
	return toReturn
}

func countCorrectPassOldRule() int {
	passs := parsePassLinesToStruct()
	counter := 0
	for _, pass := range passs {
		if pass.isPassOkOldRule() {
			counter++
		}
	}
	fmt.Println(counter)
	return counter
}

func countCorrectPassNewRule() int {
	passs := parsePassLinesToStruct()
	counter := 0
	for _, pass := range passs {
		if pass.isPassOkNewRule() {
			counter++
		}
	}
	fmt.Println(counter)
	return counter
}
