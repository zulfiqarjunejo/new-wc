package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"os"
	"unicode"
)

func getCountsOf(filename string) (Counts, error) {
	var counts Counts

	file, err := os.Open(filename)
	if err != nil {
		log.Printf("getCountsOf(%s) failed to open file\n", filename)
		return counts, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	var previousCharacterAsRune rune
	for {
		characterAsRune, size, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				if previousCharacterAsRune != rune(0) && !unicode.IsSpace(previousCharacterAsRune) {
					counts.NumberOfWords++
				}
				break
			}
			return counts, nil
		}

		if characterAsRune == rune('\n') {
			counts.NumberOfLines++
		} else if !unicode.IsSpace(previousCharacterAsRune) && unicode.IsSpace(characterAsRune) {
			counts.NumberOfWords++
		}

		counts.NumberOfBytes += size
		counts.NumberOfCharacters++

		previousCharacterAsRune = characterAsRune
	}

	return counts, nil
}

func main() {
	var cpo CountPrinterOpts

	flag.BoolVar(&cpo.ShouldPrintNumberOfBytes, "c", false, "count number of bytes")
	flag.BoolVar(&cpo.ShouldPrintNumberOfCharacters, "m", false, "counter number of characters")
	flag.BoolVar(&cpo.ShouldPrintNumberOfLines, "l", false, "count number of lines")
	flag.BoolVar(&cpo.ShouldPrintNumberOfWords, "w", false, "count number of words")

	flag.Parse()

	filename := flag.Arg(0)

	counts, err := getCountsOf(filename)
	if err != nil {
		log.Fatalf("Could not compute counts of file %s\n", filename)
	}

	countPrinter := CountPrinter{
		Opts: cpo,
	}
	countPrinter.Print(counts, filename)
}
