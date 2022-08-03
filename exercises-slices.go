package main

import (
	"golang.org/x/tour/pic"
) 

func Pic(dx, dy int) [][]uint8 {

	// Return a slice of length dy
	// Each element in slice is a slice of dx unsigned integers
	
	// Step 1: Make an array of arrays of size dy and type uint8
	ret := make([][]uint8, dy)
	
	// Step 2: Iterate over array of arrays 
	for y := range ret {
		// Make each subarray of type uint8 and size dx
		ret[y] = make([]uint8, dx)
		// Iterate over each subarray and add a value to that index of matrix
		for x := range ret[y] {
			ret[y][x] = uint8(x+y)
		}
	}
	return ret
}

func main() {
	pic.Show(Pic)
}
