package main

import (
	"fmt"
	"sort"
)

// CopySlice copies a slice of float64s
func CopySlice(vals []int) []int {
	copiedVals := make([]int, len(vals))
	copy(copiedVals, vals)
	return copiedVals
}

// SortedCopy returns a sorted copy of ints
func SortedCopy(vals []int) (copy []int) {
	copy = CopySlice(vals)
	sort.Ints(copy)
	return copy
}

func Percentile(vals []int, percent float64) (percentile int, err error) {

	if len(vals) == 0 {
		return 0, fmt.Errorf("No values")
	}
	// Percentile must be between 0 and 100
	if percent < 0 || percent > 100 {
		return 0, fmt.Errorf("Percentile must be between 0 and 100")
	}

	// Start by sorting a copy of the slice
	c := SortedCopy(vals)

	// Have percent = 0 return the min; percent = 100 return the max
	if percent == 0 {
		return c[0], nil
	} else if percent == 100 {
		return c[len(c)-1], nil
	}

	// Multiply percent by length of input
	index := (percent / 100) * float64(len(c))

	// Check if the index is a whole number
	if index == float64(int64(index)) {

		// Convert float to int
		i := int(index)

		// Find the value at the index
		percentile = c[i-1]

	} else if index > 1 {

		// Convert float to int via truncation
		i := int(index)

		// Find the average of the index and following values
		percentile = int(float64(c[i-1] + c[i]) / 2.)

	} else {
		percentile = c[0]
	}
	return percentile, nil
}

func RemoveOutliers(s []int) []int {
	// first, remove any 0s
	var nonzeros []int
	for _, val := range s {
		if val > 0 {
			nonzeros = append(nonzeros, val)
		}
	}
	fmt.Println(nonzeros)
	return nonzeros
}

