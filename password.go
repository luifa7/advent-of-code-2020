package main

// Password is a new type
// which is a structure to collect validations
type Password struct {
	minTimes int
	maxTimes int
	letter   string
	pass     string
}
