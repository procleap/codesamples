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

	fmt.Println(f.Section(".text"))
}
