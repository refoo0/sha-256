package main

import (
	"fmt"
	"log/slog"

	"github.com/refoo0/sha-256/utils"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sha256",
	Short: "A simple SHA-256 implementation in Go",
	Long:  `This is a simple implementation of the SHA-256 hash function in Go.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		slog.Error("Error executing command", "err", err)
	}
}

func init() {
	rootCmd.AddCommand(hashCmd)
}

var hashCmd = &cobra.Command{
	Use:   "hash",
	Short: "Compute the SHA-256 hash of the input string",
	Long:  `Compute the SHA-256 hash of the input string.`,
	Args:  cobra.MinimumNArgs(1),
	Run:   hashCmdRun,
}

func hashCmdRun(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		slog.Error("No input provided")
		return
	}
	input := args[0]
	hash := Hash(input)
	fmt.Printf("Input: %s\nHash: %s\n", input, hash)
}

// Konstanten für die SHA-256 Kompressionsfunktion

var k = [64]uint32{
	0x428a2f98, 0x71374491, 0xb5c0fbcf, 0xe9b5dba5, 0x3956c25b, 0x59f111f1, 0x923f82a4, 0xab1c5ed5,
	0xd807aa98, 0x12835b01, 0x243185be, 0x550c7dc3, 0x72be5d74, 0x80deb1fe, 0x9bdc06a7, 0xc19bf174,
	0xe49b69c1, 0xefbe4786, 0x0fc19dc6, 0x240ca1cc, 0x2de92c6f, 0x4a7484aa, 0x5cb0a9dc, 0x76f988da,
	0x983e5152, 0xa831c66d, 0xb00327c8, 0xbf597fc7, 0xc6e00bf3, 0xd5a79147, 0x06ca6351, 0x14292967,
	0x27b70a85, 0x2e1b2138, 0x4d2c6dfc, 0x53380d13, 0x650a7354, 0x766a0abb, 0x81c2c92e, 0x92722c85,
	0xa2bfe8a1, 0xa81a664b, 0xc24b8b70, 0xc76c51a3, 0xd192e819, 0xd6990624, 0xf40e3585, 0x106aa070,
	0x19a4c116, 0x1e376c08, 0x2748774c, 0x34b0bcb5, 0x391c0cb3, 0x4ed8aa4a, 0x5b9cca4f, 0x682e6ff3,
	0x748f82ee, 0x78a5636f, 0x84c87814, 0x8cc70208, 0x90befffa, 0xa4506ceb, 0xbef9a3f7, 0xc67178f2,
}

func rotateRight(word uint32, n int) uint32 {
	return (word >> n) | (word << (32 - n))
}

func shiftRight(word uint32, n int) uint32 {
	return word >> n
}

func sha256(input []byte) [32]byte {
	// IV-Werte
	hashes := [8]uint32{
		0x6a09e667,
		0xbb67ae85,
		0x3c6ef372,
		0xa54ff53a,
		0x510e527f,
		0x9b05688c,
		0x1f83d9ab,
		0x5be0cd19,
	}

	paddedMessage := utils.PadMessage(input)

	for n := 0; n < len(paddedMessage); n += 64 {
		var w [64]uint32

		// Nachricht in Block kopieren und in Big-Endian umwandeln
		block := paddedMessage[n : n+64]
		for i := 0; i < 16; i++ {
			w[i] = uint32(block[i*4]) << 24
			w[i] += uint32(block[i*4+1]) << 16
			w[i] += uint32(block[i*4+2]) << 8
			w[i] += uint32(block[i*4+3])
		}

		// Nachrichtenblock erweitern
		for i := 16; i < 64; i++ {
			s0 := rotateRight(w[i-15], 7) ^ rotateRight(w[i-15], 18) ^ shiftRight(w[i-15], 3)
			s1 := rotateRight(w[i-2], 17) ^ rotateRight(w[i-2], 19) ^ shiftRight(w[i-2], 10)
			w[i] = w[i-16] + s0 + w[i-7] + s1
		}

		a, b, c, d, e, f, g, h := hashes[0], hashes[1], hashes[2], hashes[3], hashes[4], hashes[5], hashes[6], hashes[7]

		// Kompressionsfunktion
		for i := 0; i < 64; i++ {
			s1 := rotateRight(e, 6) ^ rotateRight(e, 11) ^ rotateRight(e, 25)
			ch := (e & f) ^ (^e & g)
			tmp := h + s1 + ch + k[i] + w[i]

			s0 := rotateRight(a, 2) ^ rotateRight(a, 13) ^ rotateRight(a, 22)
			maj := (a & b) ^ (a & c) ^ (b & c)

			h = g
			g = f
			f = e
			e = d + tmp
			d = c
			c = b
			b = a
			a = tmp + s0 + maj
		}

		hashes[0] += a
		hashes[1] += b
		hashes[2] += c
		hashes[3] += d
		hashes[4] += e
		hashes[5] += f
		hashes[6] += g
		hashes[7] += h
	}

	// Ergebnis von Big-Endian zu Little-Endian konvertieren
	var hash [32]byte
	for i := 0; i < 8; i++ {
		hash[i*4] = byte(hashes[i] >> 24)
		hash[i*4+1] = byte(hashes[i] >> 16)
		hash[i*4+2] = byte(hashes[i] >> 8)
		hash[i*4+3] = byte(hashes[i])
	}

	return hash
}

func main() {

	Execute()

	// Aufgabe 2 a)
	message := ""
	expectedHash := "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"
	messageHash := Hash(message)
	fmt.Printf("%s\n%s\nGleicher Hash: %t\n", messageHash, expectedHash, messageHash == expectedHash)

	message = "abc"
	expectedHash = "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad"
	messageHash = Hash(message)
	fmt.Printf("%s\n%s\nGleicher Hash: %t\n", messageHash, expectedHash, messageHash == expectedHash)

	message = "abcdbcdecdefdefgefghfghighijhijkijkljklmklmnlmnomnopnopq"
	expectedHash = "248d6a61d20638b8e5c026930c3e6039a33ce45964ff2167f6ecedd419db06c1"
	messageHash = Hash(message)
	fmt.Printf("%s\n%s\nGleicher Hash: %t\n", messageHash, expectedHash, messageHash == expectedHash)

	// Aufgabe 2 b)
	//    |XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX|                               |XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX|                               |XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX|                               |XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX|                               |XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX|                               |XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX|                               |XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX|                               |
	m1 := "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	m2 := "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	m1Hash := Hashb(m1)
	m2Hash := Hashb(m2)
	fmt.Printf("%s\n%s\nGleicher Hash: %t\n", m1Hash, m2Hash, m1Hash == m2Hash)

	m1 = "Aufgabe 2 (Implementierung und Analyse des SHA-256)"
	m2 = "Aufgabe 2 (Implementierung und Analyse des SHA-224)"
	m1Hash = Hashb(m1)
	m2Hash = Hashb(m2)
	fmt.Printf("%s\n%s\nGleicher Hash: %t\n", m1Hash, m2Hash, m1Hash == m2Hash)

	m1 = "Implementieren Sie in einer Programmiersprache Ihrer Wahl die kryptographische Hash-Funktion\nSHA-256. Falls Sie sich für die Programmiersprache C entscheiden, können Sie die zur Verfügung\ngestellte teilweise Implementierung ergänzen."
	m2 = "Implementieren Sie in einer Programmiersprache Ihrer Wahl die kryptographische Hash-Funktion\nSHA-224. Falls Sie sich für die Programmiersprache C entscheiden, können Sie die zur Verfügung\ngestellte teilweise Implementierung ergänzen."
	m1Hash = Hashb(m1)
	m2Hash = Hashb(m2)
	fmt.Printf("%s\n%s\nGleicher Hash: %t\n", m1Hash, m2Hash, m1Hash == m2Hash)

	m1 = "Schwächen Sie nun Ihre Implementierung des SHA-256 wie folgt ab:\nSetzen Sie die Ausgabe der vier logischen Funktionen Ch, Maj, Sigma0 und Sigma1 auf 0, das heißt, die\nWirkung der Funktionen wird praktisch aufgehoben.\nReduzieren Sie die Anzahl der Runden innerhalb der Kompressionsfunktion von ursprünglich 64 auf 8\nRunden.\n\nFinden Sie dann eine Kollision für die so geschwächte Version des SHA-256."
	m2 = "Schwächen Sie nun Ihre Implementierung des SHA-224 wie folgt ab:\nSetzen Sie die Ausgabe der vier logischen Funktionen Ch, Maj, Sigma0 und Sigma1 auf 0, das heißt, die\nWirkung der Funktionen wird praktisch aufgehoben.\nReduzieren Sie die Anzahl der Runden innerhalb der Kompressionsfunktion von ursprünglich 64 auf 2\nRunden.\n\nFinden Sie dann eine Kollision für die so geschwächte Version des SHA-256."
	m1Hash = Hashb(m1)
	m2Hash = Hashb(m2)
	fmt.Printf("%s\n%s\nGleicher Hash: %t\n", m1Hash, m2Hash, m1Hash == m2Hash)

}

func Hash(s string) string {
	result := ""

	input := []byte(s)
	hash := sha256(input)

	for _, b := range hash {
		result += fmt.Sprintf("%02x", b)
	}

	return result
}

/************************************************
Ab hier Aufgabe 2 b)
************************************************/

func sha256b(input []byte) [32]byte {
	hashes := [8]uint32{
		0x6a09e667,
		0xbb67ae85,
		0x3c6ef372,
		0xa54ff53a,
		0x510e527f,
		0x9b05688c,
		0x1f83d9ab,
		0x5be0cd19,
	}

	paddedMessage := utils.PadMessage(input)

	for n := 0; n < len(paddedMessage); n += 64 {
		var w [64]uint32

		block := paddedMessage[n : n+64]
		for i := 0; i < 16; i++ {
			w[i] = uint32(block[i*4]) << 24
			w[i] += uint32(block[i*4+1]) << 16
			w[i] += uint32(block[i*4+2]) << 8
			w[i] += uint32(block[i*4+3])
		}
		for i := 16; i < 64; i++ {
			s0 := rotateRight(w[i-15], 7) ^ rotateRight(w[i-15], 18) ^ shiftRight(w[i-15], 3)
			s1 := rotateRight(w[i-2], 17) ^ rotateRight(w[i-2], 19) ^ shiftRight(w[i-2], 10)
			w[i] = w[i-16] + s0 + w[i-7] + s1
		}

		a, b, c, d, e, f, g, h := hashes[0], hashes[1], hashes[2], hashes[3], hashes[4], hashes[5], hashes[6], hashes[7]

		for i := 0; i < 8; i++ {
			tmp := h + k[i] + w[i]

			h = g
			g = f
			f = e
			e = d + tmp
			d = c
			c = b
			b = a
			a = tmp
		}

		hashes[0] += a
		hashes[1] += b
		hashes[2] += c
		hashes[3] += d
		hashes[4] += e
		hashes[5] += f
		hashes[6] += g
		hashes[7] += h
	}

	var hash [32]byte
	for i := 0; i < 8; i++ {
		hash[i*4] = byte(hashes[i] >> 24)
		hash[i*4+1] = byte(hashes[i] >> 16)
		hash[i*4+2] = byte(hashes[i] >> 8)
		hash[i*4+3] = byte(hashes[i])
	}

	return hash
}

func Hashb(s string) string {
	result := ""
	input := []byte(s)
	hash := sha256b(input)

	for _, b := range hash {
		result += fmt.Sprintf("%02x", b)
	}

	return result
}
