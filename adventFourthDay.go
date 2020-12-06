package main

import (
	"fmt"
	"strconv"
	"strings"
)

var pathToInputFourthDay = pwd + "/inputs/input_04"

func countCorrectPassports() int {
	pasportsInStrings := readFileLines(pathToInputFourthDay)

	correctPassportCounter := 0
	startSubSlice := 0
	endSubSlice := 0

	for i, s := range pasportsInStrings {
		if s == "" {
			endSubSlice = i
			if analyzePassportWithValidations(pasportsInStrings[startSubSlice:endSubSlice]) {
				correctPassportCounter++
			}
			startSubSlice = i
		}
	}
	if analyzePassportWithValidations(pasportsInStrings[startSubSlice:]) {
		correctPassportCounter++
	}
	fmt.Println(correctPassportCounter)
	return correctPassportCounter
}

func analyzePassportSimpleRule(pasportsInSlice []string) bool {
	isCidIn := false
	var sliceWithAllValues []string
	for _, s := range pasportsInSlice {
		subPass := strings.Fields(s)
		sliceWithAllValues = append(sliceWithAllValues, subPass...)
	}
	for _, s := range sliceWithAllValues {
		fmt.Println(s, " size ", len(sliceWithAllValues))
		if strings.Contains(s, "cid") {
			isCidIn = true
		}
	}
	if (len(sliceWithAllValues) == 7 && !(isCidIn)) || (len(sliceWithAllValues) == 8) {
		fmt.Println("TRUE")
		return true
	}
	return false
}

func analyzePassportWithValidations(pasportsInSlice []string) bool {
	isCidIn := false
	var sliceWithAllValues []string
	for _, s := range pasportsInSlice {
		subPass := strings.Fields(s)
		sliceWithAllValues = append(sliceWithAllValues, subPass...)
	}
	for _, s := range sliceWithAllValues {
		fmt.Println(s, " size ", len(sliceWithAllValues))
		if strings.Contains(s, "cid") {
			isCidIn = true
		}
	}

	if (len(sliceWithAllValues) == 7 && !(isCidIn)) || (len(sliceWithAllValues) == 8) {
		newPassport := passport{"", "", "", "", -1, -1, -1, -1}
		for _, s := range sliceWithAllValues {
			keyValue := strings.Split(s, ":")
			switch keyValue[0] {
			case "ecl":
				newPassport.ecl = keyValue[1]
			case "hcl":
				newPassport.hcl = keyValue[1]
			case "hgt":
				newPassport.hgt = keyValue[1]
			case "pid":
				newPassport.pid = keyValue[1]
			case "eyr":
				newPassport.eyr, _ = strconv.Atoi(keyValue[1])
			case "iyr":
				newPassport.iyr, _ = strconv.Atoi(keyValue[1])
			case "cid":
				newPassport.cid, _ = strconv.Atoi(keyValue[1])
			case "byr":
				newPassport.byr, _ = strconv.Atoi(keyValue[1])

			}
		}
		if newPassport.isPassportValid() {
			fmt.Println("TRUE")
			return true
		}
	}
	return false
}
