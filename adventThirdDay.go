package main

import "fmt"

var pathToInputThirdDay = pathToInputs + "03"

func reedTheMap() {
	results := readFileLines(pathToInputThirdDay)
	lenOfLine := len(results[0])
	lenOfResults := len(results)

	startRight := 0
	startDown := 0

	trees := 0
	for startDown < lenOfResults-1 {
		startDown++
		startRight += 3
		if startRight >= lenOfLine {
			startRight = startRight - lenOfLine
		}

		if string(results[startDown][startRight]) == "#" {
			trees++
		}

		fmt.Println("trees: ", trees, " on ", startDown, " line ", startRight)

	}

}

func reedTheMapPart2() {
	lines := readFileLines(pathToInputThirdDay)
	trees := travel(lines, 1, 1)
	trees = trees * travel(lines, 3, 1)
	trees = trees * travel(lines, 5, 1)
	trees = trees * travel(lines, 7, 1)
	trees = trees * travel(lines, 1, 2)
	fmt.Println(trees)

}

func travel(lines []string, right int, down int) int {
	startRight := 0
	trees := 0

	i := down
	for i < len(lines) {
		startRight += right
		if startRight > len(lines[i])-1 {
			startRight = startRight - len(lines[i])
		}
		if lines[i][startRight] == '#' {
			trees++
		}

		i += down
	}
	fmt.Println(trees)
	return trees
}
