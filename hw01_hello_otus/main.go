package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	stringToReverse := "Hello, OTUS!"
	reversedString := stringutil.Reverse(stringToReverse)
	fmt.Println(reversedString)
}
