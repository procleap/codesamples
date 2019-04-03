package main

// Processor modes: either 32-bit or 64-bit.
const (
	mode32 = 32
	mode64 = 64
)

// Disasm represents an disassembled program in memory.
type Disasm struct {
	text  []byte
	ib    uint64
	entry uint64
	mode  uint64
}
