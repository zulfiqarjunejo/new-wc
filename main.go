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

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to open %s: %s\n", filename, err.Error())
	}
	defer file.Close()

	content := make([]byte, 9999999)
	n, err := file.Read(content)

	fmt.Printf("%d %s\n", n, filename)
}
