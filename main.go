package main

import (
	"flag"
	"fmt"
)

func main() {

	dspathPtr := flag.String("dspath", "", "data set directory path")
	outPtr := flag.String("out", "", "output directory path")
	flag.Int("create_subdir", 0, "dummy flag for compatibility")
	platePtr := flag.String("plate", "", "plate name")

	flag.Parse()

	fmt.Println("ds dir:", *dspathPtr)
	fmt.Println("out dir:", *outPtr)
	fmt.Println("plate name:", *platePtr)
}
