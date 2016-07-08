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
	elizaName := "Eliza"
	niallName := "Niall"

	niall.Init()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		niall.Learn(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	elizaSaysHi := goeliza.ElizaHi()
	fmt.Println(elizaName + ": " + elizaSaysHi)
	niall.Learn(elizaSaysHi)

	// Number of lines written by Niall
	numlines := 10
	if len(os.Args) > 1 {
		numlinesString := os.Args[1]
		if numlinesInt, err := strconv.Atoi(numlinesString); err == nil {
			numlines = numlinesInt
		}
	}

	for x := 0; x < numlines; x++ {
		statement := niall.Talk()
		fmt.Printf("%s: %s\n", niallName, statement)
		fmt.Printf("%s: %s\n", elizaName, goeliza.ReplyTo(statement))
	}
}
