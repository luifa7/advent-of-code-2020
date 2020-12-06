package main

import (
	"fmt"
	"strings"
)

var pathToInputSixthDay = pathToInputs + "06"

func sumOfNumberOFQuestionsWithYes() int {
	answersInSlice := getSliceOfSlicesWithoutBlankSeparation(readFileLines(pathToInputSixthDay))
	yesAnswersCounter := 0
	for _, a := range answersInSlice {
		yesAnswersCounter += analyzeAnswersSecondRule(a)
	}
	fmt.Println(yesAnswersCounter)
	return yesAnswersCounter
}

func analyzeAnswersFirstRule(answersInSlice []string) int {
	return len(removeDuplicateValues(getOneStringWithAllValuesInSlice(answersInSlice)))
}

func analyzeAnswersSecondRule(answersInSlice []string) int {
	amountOfPeopleAnswering := len(answersInSlice)
	oneBigString := getOneStringWithAllValuesInSlice(answersInSlice)
	toReturn := 0
	for _, letter := range removeDuplicateValues(oneBigString) {
		timesAnswered := strings.Count(oneBigString, string(letter))
		fmt.Println("Here is how many times", string(letter), "was answered", timesAnswered, "for people", amountOfPeopleAnswering)
		if timesAnswered == amountOfPeopleAnswering {
			toReturn++
		}
	}
	return toReturn
}
