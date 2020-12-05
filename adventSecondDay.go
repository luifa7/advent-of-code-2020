package main

import (
	"fmt"
	"strconv"
	"strings"
)

var pathToInputSecondDay = "/Users/lgodoyalvarez/other_projects/gocode/advent/input_02"

func parsePassLinesToStruct() []Password {
	resutl := readFileLines(pathToInputSecondDay)
	var toReturn []Password
	for _, r := range resutl {
		s := strings.Fields(r)
		rangeForLetter := strings.Split(s[0], "-")
		min, _ := strconv.Atoi(rangeForLetter[0])
		max, _ := strconv.Atoi(rangeForLetter[1])
		newPass := Password{min, max, string(s[1][0]), s[2]}
		toReturn = append(toReturn, newPass)
	}
	return toReturn
}

func isPassOkOldRule(pass Password) bool {
	timesLetterInPass := strings.Count(pass.pass, pass.letter)
	if pass.minTimes <= timesLetterInPass && pass.maxTimes >= timesLetterInPass {
		return true
	}
	return false
}

func isPassOkNewRule(pass Password) bool {
	letterOnMinPossition := string(pass.pass[pass.minTimes-1])
	letterOnMaxPossition := string(pass.pass[pass.maxTimes-1])

	if letterOnMinPossition == pass.letter {
		if letterOnMaxPossition == pass.letter {
			return false
		}
		return true
	} else if letterOnMaxPossition == pass.letter {
		return true
	}

	return false
}

func countCorrectPassOldRule() int {
	passs := parsePassLinesToStruct()
	counter := 0
	for _, pass := range passs {
		if isPassOkOldRule(pass) {
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
		if isPassOkNewRule(pass) {
			counter++
		}
	}
	fmt.Println(counter)
	return counter
}
