package main

import (
	"debug/pe"
	"fmt"
	"log"
	"os"
	"sort"

	"golang.org/x/arch/x86/x86asm"
)

// Processor modes: either 32 or 64-bit.
const (
	mode32 = 32
	mode64 = 64
)

func main() {
	f, err := pe.Open(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	code, _ := f.Section(".text").Data()
	Disasm(code, mode32)
}

// Disasm disassembles either 32 or 64-bit intel instructions.
func Disasm(data []byte, mode uint64) {
	disasList := make(map[int]string)
	p := 0

	for p < len(data) {
		op, _ := x86asm.Decode(data[p:], int(mode))
		x86asm.IntelSyntax(op, mode, nil)
		disasList[p] = x86asm.IntelSyntax(op, mode, nil)
		p += op.Len
	}

	Print(disasList)
}

func Print(m map[int]string) {
	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		fmt.Printf("%8x:    %s\n", k, m[k])
	}
}
