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

	// Loop trough all PE sections.
	sep := strings.Repeat("-", 17)
	fmt.Println(sep)
	for _, section := range f.Sections {
		fmt.Printf("Section: %s\n", section.Name)
		fmt.Printf("Size: %d bytes\n", section.Size)
		Dig(section, caveSize)
		fmt.Println(sep)
	}
}

// Dig searches PE section for a code cave that is at least n bytes in size.
func Dig(s *pe.Section, n int) {
	data, err := s.Data()
	if err != nil {
		log.Fatalln(err)
	}

	// Start digging...
	var index, begin, end, count int
	for i, v := range data {
		switch {
		case v == 0:
			count++
		case count >= n:
			// Cave found!
			index++
			begin = i - count
			end = i
			fmt.Printf("# Cave %d\n", index)
			fmt.Printf("\tSize         : %d bytes\n", count)
			fmt.Printf("\tStart Address: %#x\n", begin)
			fmt.Printf("\tEnd Address  : %#x\n", end)
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
