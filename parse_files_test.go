package main

import (
	"fmt"
	"testing"
)

// Test for panics based on: https://stackoverflow.com/questions/31595791/how-to-test-panics
func TestParseBarcodeFIfile(t *testing.T) {
	tables := []struct{
		f string // name of the file
	}{
		{"./data/DPK.CP001_A549_24H_X1_B42/DPK.CP001_A549_24H_X1_B42_E14.txt",
			},
	}

	for _, tt := range tables {
		res, err := ParseBarcodeFIfile(tt.f)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(res)
	}
}
