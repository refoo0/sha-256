package core

import (
	"fmt"
)

// Padding the message
// [Message] [1 Bit] [0 Bits] [Length of Message in Bits (64 Bit)]
func PadMessage(message []byte, verbose bool) []byte {

	if verbose {
		fmt.Println("Message in bits:\n" + BytesToBits(message) + "\n")

		fmt.Println("--- Padding the message ---")
		fmt.Println("Add '1' bit , then '0' bits until message length is 448 mod 512 (56 mod 64 bytes)")
		fmt.Println("Finally, append the original message length as a 64-bit")
	}
	originalByteLen := uint64(len(message))
	originalBitLen := originalByteLen * 8
	message = append(message, 0x80)

	for len(message)%64 != 56 {
		message = append(message, 0x00)
	}

	for i := 7; i >= 0; i-- {
		message = append(message, byte(originalBitLen>>(uint(i)*8)))
	}
	if verbose {
		fmt.Println("Message in bits after appending original length:\n" + BytesToBits(message))
	}
	return message
}

func Sigma0(x uint32) uint32 {
	return RotateRight(x, 7) ^ RotateRight(x, 18) ^ ShiftRight(x, 3)
}

func Sigma1(x uint32) uint32 {
	return RotateRight(x, 17) ^ RotateRight(x, 19) ^ ShiftRight(x, 10)
}

func BigSigma1(x uint32) uint32 {
	return RotateRight(x, 6) ^ RotateRight(x, 11) ^ RotateRight(x, 25)
}

func BigSigma0(x uint32) uint32 {
	return RotateRight(x, 2) ^ RotateRight(x, 13) ^ RotateRight(x, 22)
}
