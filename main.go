package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//Calling the count function to count the number of words
	//received ferom the standard input and printing it out
	fmt.Println(count(os.Stdin))
}

func count(r io.Reader) int {
	// a scanner is used to read a text from a reader (such as files)
	scanner := bufio.NewScanner(r)

	// define the scanner split type to words (default is split by line)
	scanner.Split(bufio.ScanWords)

	// define a counter
	wc := 0

	// for every word scanned, increment the counter
	for scanner.Scan() {
		wc++
	} 

	//return the total
	return wc
}