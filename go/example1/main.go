package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	lines := flag.Bool("l", false, "Count lines")
	bytes := flag.Bool("b", false, "Count bytes")

	flag.Parse()

	fmt.Println("Flag #1 is ", *lines)
	fmt.Println("Flag #2 is ", *bytes)

	fmt.Println(count(os.Stdin, *lines, *bytes))
}

func count(r io.Reader, countlines bool, countBytes bool) int {
	scanner := bufio.NewScanner(r)

	if countBytes {
		scanner.Split(bufio.ScanBytes)
	} else {
		if !countlines {
			scanner.Split(bufio.ScanWords)
		} else {
			scanner.Split(bufio.ScanLines)
		}
	}

	wc := 0

	for scanner.Scan() {
		wc++
	}

	return wc
}
