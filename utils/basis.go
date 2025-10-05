package utils

import (
	"log/slog"
)

// Padding der Nachricht gemäß SHA-256 Spezifikation
// [Nachricht] [100000...0] [Länge der Nachricht in Bits (64 Bit)]
func PadMessage(message []byte) []byte {

	slog.Info("Message in bits:\n" + BytesToBits(message))

	originalByteLen := uint64(len(message))
	originalBitLen := originalByteLen * 8
	message = append(message, 0x80)

	slog.Info("Message in bits after appending 1 bit:\n" + BytesToBits(message))

	for len(message)%64 != 56 {
		message = append(message, 0x00)
	}
	slog.Info("Message in bits after padding with 0s:\n" + BytesToBits(message))
	for i := 7; i >= 0; i-- {
		message = append(message, byte(originalBitLen>>(uint(i)*8)))
	}
	slog.Info("Message in bits after appending original length:\n" + BytesToBits(message))
	return message
}
