package main

import (
	"debug/pe"
	"fmt"
	"log"
	"os"

	"golang.org/x/arch/x86/x86asm"
)

const (
	bit32 = 32
	bit64 = 64
)

func main() {
	f, err := pe.Open(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	code, _ := f.Section(".text").Data()
	Disasm(code, bit32)
}

// Disasm disassembles either 32 or 64-bit intel instructions.
func Disasm(data []byte, bit uint64) {
	p := 0
	for p < len(data) {
		op, _ := x86asm.Decode(data[p:], int(bit))
		x86asm.IntelSyntax(op, bit, nil)
		fmt.Println(x86asm.IntelSyntax(op, bit, nil))
		p += op.Len
	}
}
