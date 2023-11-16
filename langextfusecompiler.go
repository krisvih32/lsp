package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("in main")
	/*os := */ getinfo(len(os.Args), os.Args)
	infilename, outfilename := getfilenames(os.Args)
	infilecontents, err := os.ReadFile(infilename)
	handleerror(err)
	os.WriteFile(outfilename, compile(infilecontents), 1777)

	return
}
