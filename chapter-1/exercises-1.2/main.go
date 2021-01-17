package main

import (
	"fmt"
	"os"
)

//Exercise 1.2: Modify the echo program to print the index and value of each of its arguments, one per line.

func main() {
	// fmt.Println(strings.Join(os.Args[1:], " ")) //Original
	for i, arg := range os.Args {
		fmt.Printf("[%d]: %s\n", i, arg)
	}
}
