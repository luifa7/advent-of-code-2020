package main

import (
	"fmt"
	"regexp"
	"strconv"
)

// passport is a new type
// which is a structure to collect passport information
type passport struct {
	ecl string
	hcl string
	hgt string
	pid string
	eyr int
	iyr int
	byr int
	cid int // optional
}

func (p passport) isByrValid() bool {
	return inBetween(p.byr, 1920, 2002)
}
func (p passport) isIyrValid() bool {
	return inBetween(p.iyr, 2010, 2020)
}
func (p passport) isEyrValid() bool {
	return inBetween(p.eyr, 2020, 2030)
}
func (p passport) isHgtValid() bool {
	if inBetween(len(p.hgt), 4, 5) {
		var meassureType string
		var value int
		if len(p.hgt) == 5 {
			meassureType = p.hgt[3:]
			value, _ = strconv.Atoi(p.hgt[0:3])
		} else {
			meassureType = p.hgt[2:]
			value, _ = strconv.Atoi(p.hgt[0:2])
		}
		if meassureType == "cm" {
			return inBetween(value, 150, 193)
		}
		if meassureType == "in" {
			return inBetween(value, 59, 76)
		}
	}
	return false
}
func (p passport) isHclValid() bool {
	if len(p.hcl) > 0 {
		if string(p.hcl[0:1]) == "#" && len(p.hcl[1:]) == 6 {
			match, _ := regexp.MatchString("([a-f0-9]+)", p.hcl[1:])
			if match {
				return true
			}
		}
	}
	return false
}
func (p passport) isEclValid() bool {
	eyeColours := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	return stringInSlice(p.ecl, eyeColours)
}
func (p passport) isPidValid() bool {
	return (isStringInt(p.pid) && len(p.pid) == 9)
}
func (p passport) isCidValid() bool {
	return true
}

func (p passport) isPassportValid() bool {
	fmt.Println(p.isByrValid(), ":", p.byr)
	fmt.Println(p.isEyrValid(), ":", p.eyr)
	fmt.Println(p.isHgtValid(), ":", p.hgt)
	fmt.Println(p.isIyrValid(), ":", p.iyr)
	fmt.Println(p.isHclValid(), ":", p.hcl)
	fmt.Println(p.isEclValid(), ":", p.ecl)
	fmt.Println(p.isPidValid(), ":", p.pid)
	fmt.Println(p.isCidValid(), ":", p.cid)
	if p.isByrValid() && p.isIyrValid() && p.isEyrValid() && p.isHgtValid() && p.isHclValid() && p.isEclValid() && p.isPidValid() && p.isCidValid() {
		return true
	}
	return false
}
