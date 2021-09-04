// match tool checks a string against a pattern.
// If it matches - prints the string, otherwise prints nothing.
package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	src, err := readInput()
	if err != nil {
		fail(err)
	}
	words := countWords(src)
	fmt.Println(words)
	fmt.Scanf("h")
}

// match returns true if src matches pattern,
// false otherwise.
func countWords(src string) int {
	buf := strings.Fields(src)
	return len(buf)
}

// readInput reads pattern and source string
// from command line arguments and returns them.
func readInput() (src string, err error) {
	reader := bufio.NewReader(os.Stdin)
	src, err = reader.ReadString('\n')
	if err != nil {
		fail(err)
	}
	if src == "" {
		return src, errors.New("missing string")
	}
	return src, nil
}

// fail prints the error and exits.
func fail(err error) {
	fmt.Println("count:", err)
	os.Exit(1)
}
