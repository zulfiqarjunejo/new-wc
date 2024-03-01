package main

import "fmt"

type CountPrinterOpts struct {
	ShouldPrintNumberOfBytes      bool
	ShouldPrintNumberOfCharacters bool
	ShouldPrintNumberOfLines      bool
	ShouldPrintNumberOfWords      bool
}

type CountPrinter struct {
	Opts CountPrinterOpts
}

type Counts struct {
	NumberOfBytes      int
	NumberOfCharacters int
	NumberOfLines      int
	NumberOfWords      int
}

func (cp CountPrinter) Print(c Counts, filename string) {
	if cp.Opts.ShouldPrintNumberOfBytes {
		fmt.Printf("  %d %s\n", c.NumberOfBytes, filename)
	} else if cp.Opts.ShouldPrintNumberOfCharacters {
		fmt.Printf("  %d %s\n", c.NumberOfCharacters, filename)
	} else if cp.Opts.ShouldPrintNumberOfLines {
		fmt.Printf("  %d %s\n", c.NumberOfLines, filename)
	} else if cp.Opts.ShouldPrintNumberOfWords {
		fmt.Printf("  %d %s\n", c.NumberOfWords, filename)
	} else {
		fmt.Printf("\n  %d %d %d %s\n", c.NumberOfLines, c.NumberOfWords, c.NumberOfBytes, filename)
	}
}
