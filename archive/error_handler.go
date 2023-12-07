package main

import "fmt"

func handle_error(err []error) {
	if err != nil {
		panic(fmt.Sprintf(string(err)))
	}

}
