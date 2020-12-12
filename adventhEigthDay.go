package main

import (
	"fmt"
	"strconv"
	"strings"
)

var pathToInputEigthDay = pathToInputs + "08"

func breakAccBeforeRepeat() {
	control := make(map[int]bool)
	instructions := readFileLines(pathToInputEigthDay)
	acc := 0

	for i := 0; i < len(instructions); i++ {
		if control[i] {
			break
		}
		control[i] = true
		v := instructions[i]
		c := strings.Split(v, " ")
		n, _ := strconv.Atoi(c[1][1:])

		switch c[0] + string(c[1][0]) {
		case "acc+":
			acc += n
		case "acc-":
			acc -= n
		case "jmp+":
			i += n
			i--
		case "jmp-":
			i -= n
			i--
		}
	}
	fmt.Println(acc)
}
