package main

import (
	"debug/pe"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	// Validate command line parameters.
	if len(os.Args) != 3 {
		Usage()
	}

	// Check if we got the correct number for cave size.
	caveSize, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalln(err)
	} else if caveSize <= 0 {
		fmt.Println("Error: <cave size> must be bigger than zero")
		Usage()
	}

	// Try to open input file or exit program otherwise.
	f, err := pe.Open(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	// Loop through all PE sections.
	sep := strings.Repeat("-", 35)
	fmt.Println(sep)
	for _, s := range f.Sections {
		fmt.Printf("Section %s (%d bytes)\n", s.Name, s.Size)
		Dig(s, caveSize)
		fmt.Println(sep)
	}
}

// Dig searches PE section for a code cave that is at least n bytes in size.
func Dig(s *pe.Section, n int) {
	data, err := s.Data()
	if err != nil {
		log.Fatalln(err)
	}
	data = append(data, 0xff) // Sentinel.

	// Start digging...
	var index, begin, end, count int
	for i, b := range data {
		switch {
		case b == 0:
			count++
		case count >= n:
			// Cave found!
			index++
			begin = i - count
			end = i
			fmt.Printf("# Cave %d\n", index)
			fmt.Printf("\tSize        : %d bytes\n", count)
			fmt.Printf("\tOffset Start: %xh\n", uint32(begin)+s.Offset)
			fmt.Printf("\tOffset End  : %xh\n", uint32(end)+s.Offset)
			fallthrough
		default:
			count = 0 // reset counter to keep digging.
		}
	}
	if index == 0 {
		fmt.Println("Sorry, no cave found :(")
	}
}

// Usage prints the program usage to stdout and exit.
func Usage() {
	fmt.Printf("Usage:\n\t%s <file.exe> <cave size>\n", filepath.Base(os.Args[0]))
	fmt.Printf("Example:\n\t%s calc.exe 80\n", filepath.Base(os.Args[0]))
	os.Exit(1)
}
