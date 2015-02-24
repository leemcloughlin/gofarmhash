package farmhash

// platform.go based on the original C++ platform.cc

import (
	"encoding/binary"
)

// Note: I used to call binary.LittleEndian.Uint32 and Uint64 inline but it
// made comparing the code to the original much trickier

func fetch32(p []byte) uint32 {
	return binary.LittleEndian.Uint32(p)
}

func fetch64(p []byte) uint64 {
	return binary.LittleEndian.Uint64(p)
}

// rotate32 is a bitwise rotate
func rotate32(val uint32, shift uint) uint32 {
	if shift == 0 {
		return val
	}
	return val>>shift | val<<(32-shift)
}

// rotate64 is a bitwise rotate
func rotate64(val uint64, shift uint) uint64 {
	if shift == 0 {
		return val
	}
	return val>>shift | val<<(64-shift)
}
