package main

import (
	"fmt"

	"github.com/fatih/color"
)

func printerror(error_msg string) {
	color.Set(color.FgRed)
	fmt.Printf("error: ")
	color.Unset()
	fmt.Printf("%s\n", error_msg)
}
