package main

import (
	"debug/pe"
	"fmt"
	"log"
	"os"
	"sort"

	"golang.org/x/arch/x86/x86asm"
)

// Processor modes: either 32-bit or 64-bit.
const (
	mode32 = 32
	mode64 = 64
)

// Disasm represents an disassembled program in memory.
type Disasm struct {
	text       []byte
	ib         uint64
	entry      uint64
	mode       uint64
	disasmList map[uint64]string
}

// NewDisasm opens a program file and reads its content for later disassembly.
func NewDisasm(filename string) *Disasm {
	d := &Disasm{disasmList: make(map[uint64]string)}
	if err := d.init(filename); err != nil {
		log.Fatalln(err)
	}
	return d
}

// init initializes Disasm struct with the filename contents.
func (d *Disasm) init(filename string) error {
	f, err := pe.Open(os.Args[1])
	if err != nil {
		return err
	}
	defer f.Close()

	// Get code (for PE file it's in .text section).
	d.text, err = f.Section(".text").Data()
	if err != nil {
		return err
	}

	// Fill out remaining Disasm fields taking into consideration if it is
	// either a 32 or 64 bit binary.
	switch oh := f.OptionalHeader.(type) {
	case *pe.OptionalHeader32:
		d.ib = uint64(oh.ImageBase)
		d.entry = d.ib + uint64(oh.AddressOfEntryPoint)
		d.mode = mode32
	case *pe.OptionalHeader64:
		d.ib = oh.ImageBase
		d.entry = d.ib + uint64(oh.AddressOfEntryPoint)
		d.mode = mode64
	}

	return nil
}

// Disasm does a linear disassembly of binary.
func (d *Disasm) Disasm() {
	var p int
	for p < len(d.text) {
		op, _ := x86asm.Decode(d.text[p:], int(d.mode))
		x86asm.IntelSyntax(op, d.mode, nil)
		d.disasmList[uint64(p)+d.ib] = x86asm.IntelSyntax(op, d.mode, nil)
		p += op.Len
	}
	d.print()
}

// print prints disassembly list to standard output.
func (d *Disasm) print() {
	var keys []uint64
	for k := range d.disasmList {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
	for _, k := range keys {
		fmt.Printf("%x: %s\n", k, d.disasmList[k])
	}
}
