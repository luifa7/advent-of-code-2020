package main

import "fmt"

func understandTheGame() {
	m := make(map[int]int)
	m[1] = 1
	m[17] = 2
	m[0] = 3
	m[10] = 4
	m[18] = 5
	m[11] = 6
	m[6] = 7

	actualPosition := 8
	lastSpokenNumber := 6
	nextNumber := 0
	for actualPosition <= 2020 {

		if val, ok := m[nextNumber]; ok {
			lastSpokenNumber = nextNumber
			m[nextNumber] = actualPosition
			nextNumber = actualPosition - val
			actualPosition++
		} else {
			lastSpokenNumber = nextNumber
			m[nextNumber] = actualPosition
			nextNumber = 0
			actualPosition++
		}

		fmt.Println("lastSpoken: ", lastSpokenNumber, "actualposition: ", actualPosition-1)

	}

}

func findTheNumber2() {
	m := make(map[int]int)
	m[1] = 1
	m[17] = 2
	m[0] = 3
	m[10] = 4
	m[18] = 5
	m[11] = 6
	m[6] = 7

	actualPosition := 8
	lastSpokenNumber := 6
	nextNumber := 0
	for actualPosition <= 30000000 {

		if val, ok := m[nextNumber]; ok {
			lastSpokenNumber = nextNumber
			m[nextNumber] = actualPosition
			nextNumber = actualPosition - val
			actualPosition++
		} else {
			lastSpokenNumber = nextNumber
			m[nextNumber] = actualPosition
			nextNumber = 0
			actualPosition++
		}
		if actualPosition%1000 == 0 || actualPosition == 30000000 || actualPosition == 30000001 {
			fmt.Println("lastSpoken: ", lastSpokenNumber, "actualposition: ", actualPosition-1)
		}

	}

}
