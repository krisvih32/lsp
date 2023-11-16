package main

import (
	"fmt"
	"runtime"
)

func getargs(argc int, argv []string) {
	fmt.Printf("in getargs")
	fmt.Printf("argc %d\n", argc)
	switch argc {
	case 1:
		{
			printerror("no input/output files inputted using args")
			fmt.Printf("case 1: no detectable errors\n")
			panic(1)
		}
	case 2:
		{
			printerror("no output file inputted using args")
			fmt.Printf("case 2: no detectable errors\n")
			panic(2)
		}
	case 3:
		{
			return
		}
	}
}

func getinfo(argc int, argv []string) (os string) {
	getargs(argc, argv)
	os = runtime.GOOS
	return
}
