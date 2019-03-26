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
		fmt.Println(sep)
	}
}
