package main

import (
	"debug/pe"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
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
		Dig(section, 150)
		fmt.Println(sep)
	}

}

// Dig searches PE section for a code cave that is at least n bytes in size.
func Dig(s *pe.Section, n int) {
	data, _ := s.Data()
	var index, begin, end, count int

	// Start digging...
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
