package main

import "os"

func main() {
	d := NewDisasm(os.Args[1])
	d.Disasm()
}
