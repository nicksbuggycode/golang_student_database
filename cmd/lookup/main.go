package main

import (
	"io"
	"nicksbuggycode/lookup"
	"os"
	"strings"
)

func GetStudent(r io.Reader) (string, error) {
	args := os.Args[1:]
	return strings.Join(args, " "), nil
}

func main() {
	// ask for input
	//make a reader to get user input

	input, _ := GetStudent(os.Stdin)
	lookup.ConnectToDBFindSome(input)
}
