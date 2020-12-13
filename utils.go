package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func readFileLines(path string) []string {
	var toReturn []string
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		toReturn = append(toReturn, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return toReturn
}

func inBetween(i, min, max int) bool {
	if (i >= min) && (i <= max) {
		return true
	}
	return false
}

func stringInSlice(s string, ss []string) bool {
	for _, b := range ss {
		if b == s {
			return true
		}
	}
	return false
}

func isStringInt(v string) bool {
	if _, err := strconv.Atoi(v); err == nil {
		return true
	}
	return false
}

func convertByteStringToInt(byteString string) int {
	i, err := strconv.ParseInt(byteString, 2, 64)
	if err != nil {
		return -1
	}
	return int(i)
}

func removeDuplicateValues(s string) string {
	var toReturn string
	for _, letter := range s {
		if !strings.Contains(toReturn, string(letter)) {
			toReturn += string(letter)
		}
	}
	return toReturn
}

func getOneStringWithAllValuesInSlice(s []string) string {
	var oneBigString string
	for _, value := range s {
		oneBigString += value
	}
	return oneBigString
}

func getSliceOfSlicesWithoutBlankSeparation(s []string) [][]string {
	var sliceToReturn [][]string
	startSubSlice := 0
	endSubSlice := 0

	for i, v := range s {
		if v == "" {
			endSubSlice = i
			sliceToReturn = append(sliceToReturn, s[startSubSlice:endSubSlice])
			startSubSlice = i + 1
		}
	}
	sliceToReturn = append(sliceToReturn, s[startSubSlice:])
	return sliceToReturn
}

func unique(strSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range strSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func convertSSIntoSI(ss []string) []int {
	var r []int

	for _, s := range ss {
		n, _ := strconv.Atoi(s)
		r = append(r, n)
	}

	return r
}
