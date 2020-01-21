package main

import (
	"bufio"
	"fmt"
	"github.com/xyproto/eliza"
	"github.com/xyproto/niall"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	elizaName := "Eliza"
	niallName := "Niall"

	niall.Init()

	// First try learning from niall.txt, if it exists.
	// If not, use stdin.
	if contents, err2 := ioutil.ReadFile("niall.txt"); err2 == nil { // no error
		niall.Learn(string(contents))
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			niall.Learn(scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
	}

	elizaSaysHi := eliza.ElizaHi()
	fmt.Println(elizaName + ": " + elizaSaysHi)
	niall.Learn(elizaSaysHi)

	// Number of lines written by Niall, by default
	numlines := 20
	if len(os.Args) > 1 {
		numlinesString := os.Args[1]
		if numlinesInt, err := strconv.Atoi(numlinesString); err == nil {
			numlines = numlinesInt
		}
	}

	var elizaSays string
	for x := 0; x < numlines; x++ {
		statement := niall.Talk()
		niall.Learn(statement)
		fmt.Printf("%s> %s\n", niallName, statement)
		elizaSays = eliza.ReplyTo(statement)
		niall.Learn(elizaSays)
		fmt.Printf("%s> %s\n", elizaName, elizaSays)
	}
}
