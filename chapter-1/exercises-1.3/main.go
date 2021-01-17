package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

/*
Exercise 1.3: Experiment to measure the difference in running time
between our potentially inefficient versions and the one that uses
strings.Join. (Section 1.6 illustrates part of the time package,
and Section 11.4 shows how to write benchmark tests for systematic
performance evaluation.)
*/

//Print elapsed time:
//fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

func main() {
	// Inefficient string
	// start := time.Now()
	// s, sep := "", ""

	// for _, arg := range os.Args[1:] {
	// 	s += sep + arg
	// 	sep = " "
	// }

	// fmt.Println(s)
	// fmt.Printf("%d nanoseconds elapsed\n", time.Since(start).Nanoseconds())

	//Efficient with string.Join()
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Printf("%d nanoseconds elapsed\n", time.Since(start).Nanoseconds())
}
