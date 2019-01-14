package main

import (
	"fmt"
	"testing"
)

// Test for panics based on: https://stackoverflow.com/questions/31595791/how-to-test-panics
func TestRemoveOutliers(t *testing.T) {
	tables := []struct{
		vals []int // input slice
	}{
		{[]int{848, 1899, 2802, 0}},
	}

	for _, tt := range tables {
		filt := RemoveOutliers(tt.vals)
		fmt.Println(filt)
	}
}




// Test for panics based on: https://stackoverflow.com/questions/31595791/how-to-test-panics
func TestPercentile(t *testing.T) {
	tables := []struct{
		vals []int // input slice
		percentile float64 // percentile value
	}{
		{[]int{10, 8, 7, 6, 9, 3, 4, 5, 1, 2},
			0},
		{[]int{10, 8, 7, 6, 9, 3, 4, 5, 1, 2},
			10},
		{[]int{10, 8, 7, 6, 9, 3, 4, 5, 1, 2},
			20},
		{[]int{10, 8, 7, 6, 9, 3, 4, 5, 1, 2},
			30},
		{[]int{10, 8, 7, 6, 9, 3, 4, 5, 1, 2},
			40},
		{[]int{10, 8, 7, 6, 9, 3, 4, 5, 1, 2},
			50},
		{[]int{10, 8, 7, 6, 9, 3, 4, 5, 1, 2},
			60},
		{[]int{10, 8, 7, 6, 9, 3, 4, 5, 1, 2},
			70},
		{[]int{10, 8, 7, 6, 9, 3, 4, 5, 1, 2},
			80},
		{[]int{10, 8, 7, 6, 9, 3, 4, 5, 1, 2},
			90},
		{[]int{10, 8, 7, 6, 9, 3, 4, 5, 1, 2},
			100},
		{[]int{10, 8, 7, 6, 9, 3, 4, 5, 1, 2},
			92},
		{[]int{10, 8, 7, 6, 9, 3, 4, 5, 1, 2},
			95},
		{[]int{10, 8, 7, 6, 9, 3, 4, 5, 1, 2},
			99},
	}

	for _, tt := range tables {
		perc, _ := Percentile(tt.vals, tt.percentile)
		fmt.Println(perc)
	}
}
