package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	_ = flag.Bool("c", false, "count number of bytes")

	flag.Parse()

	filename := flag.Arg(0)

	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read %s: %s\n", filename, err.Error())
	}

	fmt.Printf("%d %s\n", len(content), filename)
}
