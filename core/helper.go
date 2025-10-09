package core

import (
	"fmt"
)

func transformToByteArray(hashes [8]uint32) [32]byte {
	var hash [32]byte
	for i := 0; i < 8; i++ {
		hash[i*4] = byte(hashes[i] >> 24)
		hash[i*4+1] = byte(hashes[i] >> 16)
		hash[i*4+2] = byte(hashes[i] >> 8)
		hash[i*4+3] = byte(hashes[i])
	}
	return hash
}

func BytesToBits(b []byte) string {
	bits := ""
	for idx, byteVal := range b {
		for i := 7; i >= 0; i-- {
			if byteVal&(1<<i) != 0 {
				bits += "1"
			} else {
				bits += "0"
			}
		}

		// add a newline every 4 bytes
		if (idx+1)%8 == 0 {
			bits += "\n"
			continue
		}

		// add a space between bytes
		bits += " "

	}
	return bits
}

func RotateRight(word uint32, n int) uint32 {
	return (word >> n) | (word << (32 - n))
}

func ShiftRight(word uint32, n int) uint32 {
	return word >> n
}

func Ch(x, y, z uint32) uint32 {
	return (x & y) ^ (^x & z)
}
func Maj(x, y, z uint32) uint32 {
	return (x & y) ^ (x & z) ^ (y & z)
}

func ColorPrintf(color string, format string, a ...interface{}) {
	fmt.Printf(color+format+Reset, a...)
}

func ColorPrintDiff(hash1 string, hash2 string) {
	for i := range len(hash1) {
		if hash1[i] == hash2[i] {
			ColorPrintf(Green, "%c", hash1[i])
		} else {
			ColorPrintf(Red, "%c", hash1[i])
		}
	}
}
