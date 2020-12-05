package main

import "strings"

// password is a new type
// which is a structure to collect validations
type password struct {
	minTimes int
	maxTimes int
	letter   string
	pass     string
}

func (p password) isPassOkNewRule() bool {
	letterOnMinPossition := string(p.pass[p.minTimes-1])
	letterOnMaxPossition := string(p.pass[p.maxTimes-1])

	if letterOnMinPossition == p.letter {
		if letterOnMaxPossition == p.letter {
			return false
		}
		return true
	} else if letterOnMaxPossition == p.letter {
		return true
	}

	return false
}

func (p password) isPassOkOldRule() bool {
	timesLetterInPass := strings.Count(p.pass, p.letter)
	if p.minTimes <= timesLetterInPass && p.maxTimes >= timesLetterInPass {
		return true
	}
	return false
}
