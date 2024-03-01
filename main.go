package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type CountCollection struct {
	Bytes      int
	Characters int
	Lines      int
	Words      int
}

func main() {
	c := flag.Bool("c", false, "count number of bytes")
	l := flag.Bool("l", false, "count number of lines")
	w := flag.Bool("w", false, "count number of words")
	m := flag.Bool("m", false, "counter number of characters")

	flag.Parse()

	filename := flag.Arg(0)

	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read %s: %s\n", filename, err.Error())
	}

	cc := CountCollection{}

	cc.Bytes = len(content)
	cc.Characters = len(content)

	lines := strings.Split(string(content), "\n")
	cc.Lines = len(lines) - 1

	var wordCount int
	for _, line := range lines {
		wordsInLine := strings.Fields(line)
		wordCount += len(wordsInLine)
	}

	cc.Words = wordCount

	if *c {
		fmt.Printf("  %d %s\n", len(content), filename)
	} else if *l {
		lines := strings.Split(string(content), "\n")
		fmt.Printf("  %d %s\n", len(lines)-1, filename)
	} else if *w {
		var wordCount int
		lines := strings.Split(string(content), "\n")
		for _, line := range lines {
			wordsInLine := strings.Fields(line)
			wordCount += len(wordsInLine)
		}
		fmt.Printf("  %d %s\n", wordCount, filename)
	} else if *m {
		fmt.Printf("  %d %s\n", len(content), filename)
	} else {
		fmt.Printf("  %d %d %d %s\n", cc.Lines, cc.Words, cc.Bytes, filename)
	}
}
