package main

import (
	"fmt"
	"strings"
)

type rules map[string][]string

var pathToInputSeventhDay = pathToInputs + "07"

func bagsCount() {
	colours, r := getColoursAndRules(readFileLines(pathToInputSeventhDay))
	tc := "shinygold"
	rs := []string{}
	for _, c := range colours {
		if c != tc && connectionExist(c, tc, r) {
			rs = append(rs, c)
		}
	}

	fmt.Println(len(rs))
}

func connectionExist(as string, ts string, r rules) bool {
	if _, ok := r[as]; ok {
		result := false
		for _, s := range r[as] {
			if s == ts {
				return true
			}
			result = result || connectionExist(s, ts, r)
		}
		return result
	}
	return false
}

func leaveOnlyColours(ss []string) []string {
	ss = strings.Split(strings.ReplaceAll(strings.Join(ss, ""), "bag,", "bags,"), "bags,")
	for i, v := range ss {
		if i > 0 {
			ss[i] = v[1:]
		}
	}
	return ss
}

func getColoursAndRules(ss []string) ([]string, rules) {
	cs := []string{}
	rs := make(rules)
	for _, r := range ss {
		sr := strings.Split(r, " ")
		if "no" != sr[4] {
			k := sr[0] + sr[1]
			v := leaveOnlyColours(sr[5 : len(sr)-1])
			rs[k] = v
			cs = append(cs, k)
		}
	}
	return cs, rs
}
