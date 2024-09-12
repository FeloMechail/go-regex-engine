package main

import (
	"fmt"
	"os"

	"regex-engine/pkg/regex"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: regex-cli <pattern> <text>")
		os.Exit(1)
	}

	pattern := os.Args[1]
	text := os.Args[2]

	if regex.Match(pattern, text) {
		fmt.Println("Match found!")
	} else {
		fmt.Println("No match.")
	}
}
