package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
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

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
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
