package main

import (
	"debug/pe"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := pe.Open(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	// Loop trough all PE sections.
	for _, section := range f.Sections {
		fmt.Println(section)
	}
}
