// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded but its count.

// Dup2 prints the count and text of lines that appear more than
// once in the input. It reads from stdin or from a list of named files

// Dup3 prints the count and text of lines that appear more than
// once in the input. It now only reads named files, process all text at once,
// then splits the string into a slice of substrings.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()+"\t"+f.Name()]++
	}
}

func main() {
	// **********************************************************
	// DUP1
	// *** ctrl + D ends stdin stream ***
	//
	// 	counts := make(map[string]int)
	// 	input := bufio.NewScanner(os.Stdin)
	// 	for input.Scan() {
	// 		counts[input.Text()]++
	// 	}
	// 	// Note: ignoring potential errors from input.Err()
	// 	for line, n := range counts {
	// 		if n > 1 {
	// 			fmt.Printf("%d\t%s\n", n, line)
	// 		}
	// 	}
	// **********************************************************

	// **********************************************************
	// DUP2
	// counts := make(map[string]int)
	// files := os.Args[1:]
	// if len(files) == 0 {
	// 	countLines(os.Stdin, counts)
	// } else {
	// 	for _, arg := range files {
	// 		f, err := os.Open(arg)
	// 		if err != nil {
	// 			fmt.Fprintf(os.Stderr, "dupe2: %v\n", err)
	// 			continue
	// 		}
	// 		countLines(f, counts)
	// 		f.Close()
	// 	}
	// }
	// for line, n := range counts {
	// 	if n > 1 {
	// 		fmt.Printf("%d\t%s\n", n, line)
	// 	}
	// 	// Note: ignoring potential errors from input.Err()
	// }
	// **********************************************************

	// **********************************************************
	// DUP3
	//
	// counts := make(map[string]int)
	// for _, filename := range os.Args[1:] {
	// 	data, err := ioutil.ReadFile(filename)
	// 	if err != nil {
	// 		fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
	// 		continue
	// 	}
	// 	for _, line := range strings.Split(string(data), "\n") {
	// 		counts[line]++
	// 	}
	// }
	// for line, n := range counts {
	// 	if n > 1 {
	// 		fmt.Printf("%d\t%s\n", n, line)
	// 	}
	// }
	// **********************************************************

	// EXERCISE //
	// Modify Dup2 to print the names of all files in which each duplicated line occurs
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dupe2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
		// Note: ignoring potential errors from input.Err()
	}
}
