package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	//defining the boolean flag -l to count lines instead of words
	lines := flag.Bool("l", false, "Count lines")
	bytes := flag.Bool("b", false, "Count bytes")
	
	//parsing the flag provided by the user
	flag.Parse()

	//Calling the count function to count the number of words
	//received ferom the standard input and printing it out
	fmt.Println(count(os.Stdin, *lines, *bytes))
}

func count(r io.Reader, countLines bool, countBytes bool) int {
	// a scanner is used to read a text from a reader (such as files)
	scanner := bufio.NewScanner(r)

	// define the scanner split type to words or bytes(default is split by line)
	if !countLines {
		if !countBytes {
			scanner.Split(bufio.ScanWords)
		} else {
			scanner.Split(bufio.ScanBytes)
		}
		
	}

	// define a counter
	wc := 0

	// for every word scanned, increment the counter
	for scanner.Scan() {
		wc++
	} 

	//return the total
	return wc
}