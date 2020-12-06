package main

import (
	"fmt"
	"sort"
	"strings"
)

var pathToInputFifthDay = pwd + "/inputs/input_05"

func findHighestSeatID() int {
	highest := -1
	encodedSeatIds := readFileLines(pathToInputFifthDay)

	for _, encodedSeatID := range encodedSeatIds {
		encodedRow := encodedSeatID[0:7]
		encodedColumn := encodedSeatID[7:]

		binaryStringRow := strings.ReplaceAll(strings.ReplaceAll(encodedRow, "F", "0"), "B", "1")
		binaryStringColumn := strings.ReplaceAll(strings.ReplaceAll(encodedColumn, "L", "0"), "R", "1")

		row := convertByteStringToInt(binaryStringRow)
		column := convertByteStringToInt(binaryStringColumn)

		// fmt.Println(binaryStringRow, row)
		// fmt.Println(binaryStringColumn, column)

		decodedSeatID := row*8 + column

		if decodedSeatID > highest {
			highest = decodedSeatID
		}
	}

	fmt.Println(highest)
	return highest
}

func findMySeatID() int {
	var allDecodedIds []int
	encodedSeatIds := readFileLines(pathToInputFifthDay)

	for _, encodedSeatID := range encodedSeatIds {
		encodedRow := encodedSeatID[0:7]
		encodedColumn := encodedSeatID[7:]

		binaryStringRow := strings.ReplaceAll(strings.ReplaceAll(encodedRow, "F", "0"), "B", "1")
		binaryStringColumn := strings.ReplaceAll(strings.ReplaceAll(encodedColumn, "L", "0"), "R", "1")

		row := convertByteStringToInt(binaryStringRow)
		column := convertByteStringToInt(binaryStringColumn)

		decodedSeatID := row*8 + column

		allDecodedIds = append(allDecodedIds, decodedSeatID)
	}

	sort.Ints(allDecodedIds)

	mySeatID := -1
	for i, decodedID := range allDecodedIds {
		if i+1 < len(allDecodedIds) {
			if decodedID+2 == allDecodedIds[i+1] {
				mySeatID = decodedID + 1
			}
		}
	}

	fmt.Println(mySeatID)
	return mySeatID
}
