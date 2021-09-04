// match tool checks a string against a pattern.
// If it matches - prints the string, otherwise prints nothing.
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	str := getInputString()
	words := countWords(str)
	fmt.Println(words)
}

// match returns true if src matches pattern,
// false otherwise.
func countWords(src string) int {
	buf := strings.Fields(src)
	return len(buf)
}

func getInputString() string {
	input := ""
	if len(os.Args) > 1 {
		input = os.Args[1]
	}

	return input
}

// fail prints the error and exits.
func fail(err error) {
	fmt.Println("count:", err)
	os.Exit(1)
}
