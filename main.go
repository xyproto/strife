package main

import (
	"bufio"
	"fmt"
	"github.com/kennysong/goeliza"
	"github.com/xyproto/niall"
	"os"
	"strconv"
)

func main() {
	eliza_name := "Eliza"
	niall_name := "Niall"

	niall.Init()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		niall.Learn(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	eliza_says_hi := goeliza.ElizaHi()
	fmt.Println(eliza_name + ": " + eliza_says_hi)
	niall.Learn(eliza_says_hi)

	// Number of lines written by Niall
	numlines := 10
	if len(os.Args) > 1 {
		numlines_string := os.Args[1]
		if numlines_int, err := strconv.Atoi(numlines_string); err == nil {
			numlines = numlines_int
		}
	}

	for x := 0; x < numlines; x++ {
		statement := niall.Talk()
		fmt.Printf("%s: %s\n", niall_name, statement)
		fmt.Printf("%s: %s\n", eliza_name, goeliza.ReplyTo(statement))
	}
}
