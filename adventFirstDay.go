package main

import (
	"fmt"
	"strconv"
)

var pathToInputFirstDay = pathToInputs + "01"

func parsePassLinesToInt() []int {
	resutl := readFileLines(pathToInputFirstDay)
	var toReturn []int
	for _, r := range resutl {
		toInt, _ := strconv.Atoi(r)
		toReturn = append(toReturn, toInt)
	}
	return toReturn
}

func multiplyTwoThatSums2020() {
	found := false
	arrayToAnalyze := parsePassLinesToInt()
	for _, firstNumber := range arrayToAnalyze {
		for _, secondNumber := range arrayToAnalyze {
			if firstNumber+secondNumber == 2020 {
				fmt.Println(firstNumber)
				fmt.Println(secondNumber)
				fmt.Println(firstNumber * secondNumber)
				found = true
				break
			}
		}
		if found {
			break
		}
	}
}

func multiplyThreeThatSums2020() {
	found := false
	arrayToAnalyze := parsePassLinesToInt()
	for _, firstNumber := range arrayToAnalyze {
		for _, secondNumber := range arrayToAnalyze {
			for _, thirdNumber := range arrayToAnalyze {
				if firstNumber+secondNumber+thirdNumber == 2020 {
					fmt.Println(firstNumber)
					fmt.Println(secondNumber)
					fmt.Println(thirdNumber)
					fmt.Println(firstNumber * secondNumber * thirdNumber)
					found = true
					break
				}
			}
			if found {
				break
			}
		}
		if found {
			break
		}
	}
}
