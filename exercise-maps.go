package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

// Given an input string, return a map where key is word and val is number of occurrences in input string

func WordCount(s string) map[string]int {

	// Step 1: Initialize map which will be returned
	var ret = make(map[string]int)
	
	// Step 2: Iterate over string using strings.Fields which splits string by whitespace
	// We have no use for idx so can leave that blank
	for _, w:= range strings.Fields(s) {
		// Increment count of that word in map
		ret[w]++
	}
	return ret
}

func main() {
	wc.Test(WordCount)
}
