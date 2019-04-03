package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Validate command line parameters.
	if len(os.Args) != 2 {
		Usage()
	}

	// Let's go!
	d := NewDisasm(os.Args[1])
	d.Disasm()
}

// Usage prints the program usage to stdout and exit.
func Usage() {
	fmt.Printf("Usage:\n\t%s <file.exe>\n", filepath.Base(os.Args[0]))
	os.Exit(1)
}
