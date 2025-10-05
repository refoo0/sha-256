package main

import (
	"fmt"
	"log/slog"

	"github.com/refoo0/sha-256/core"
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
	message := args[0]
	var iterations int = 64
	var rounding bool = true

	if len(args) > 1 {
		// Convert iterations to int
		_, err := fmt.Sscanf(args[1], "%d", &iterations)
		if err != nil {
			slog.Error("Invalid value for iterations", "err", err)
			return
		}
		// Convert rounding to bool

		if args[2] == "true" || args[2] == "1" {
			rounding = true
		} else if args[2] == "false" || args[2] == "0" {
			rounding = false
		} else {
			slog.Error("Invalid value for rounding, must be true/false or 1/0")
			return
		}
	}

	hash := core.SHA256([]byte(message), iterations, rounding)
	fmt.Printf("Input: %s\nHash: %x\n", message, hash)
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
	hash := core.SHA256(input, 64, true)

	for _, b := range hash {
		result += fmt.Sprintf("%02x", b)
	}

	return result
}

func Hashb(s string) string {
	result := ""

	input := []byte(s)
	hash := core.SHA256(input, 8, false)

	for _, b := range hash {
		result += fmt.Sprintf("%02x", b)
	}

	return result
}
