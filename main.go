package main

import (
	"fmt"

	"github.com/refoo0/sha-256/core"
)

func main() {

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
	m1Hash := Hash8(m1)
	m2Hash := Hash8(m2)
	fmt.Printf("%s\n%s\nGleicher Hash: %t\n", m1Hash, m2Hash, m1Hash == m2Hash)

	m1 = "Aufgabe 2 (Implementierung und Analyse des SHA-256)"
	m2 = "Aufgabe 2 (Implementierung und Analyse des SHA-224)"
	m1Hash = Hash8(m1)
	m2Hash = Hash8(m2)
	fmt.Printf("%s\n%s\nGleicher Hash: %t\n", m1Hash, m2Hash, m1Hash == m2Hash)

	m1 = "Implementieren Sie in einer Programmiersprache Ihrer Wahl die kryptographische Hash-Funktion\nSHA-256. Falls Sie sich für die Programmiersprache C entscheiden, können Sie die zur Verfügung\ngestellte teilweise Implementierung ergänzen."
	m2 = "Implementieren Sie in einer Programmiersprache Ihrer Wahl die kryptographische Hash-Funktion\nSHA-224. Falls Sie sich für die Programmiersprache C entscheiden, können Sie die zur Verfügung\ngestellte teilweise Implementierung ergänzen."
	m1Hash = Hash8(m1)
	m2Hash = Hash8(m2)
	fmt.Printf("%s\n%s\nGleicher Hash: %t\n", m1Hash, m2Hash, m1Hash == m2Hash)

	m1 = "Schwächen Sie nun Ihre Implementierung des SHA-256 wie folgt ab:\nSetzen Sie die Ausgabe der vier logischen Funktionen Ch, Maj, Sigma0 und Sigma1 auf 0, das heißt, die\nWirkung der Funktionen wird praktisch aufgehoben.\nReduzieren Sie die Anzahl der Runden innerhalb der Kompressionsfunktion von ursprünglich 64 auf 8\nRunden.\n\nFinden Sie dann eine Kollision für die so geschwächte Version des SHA-256."
	m2 = "Schwächen Sie nun Ihre Implementierung des SHA-224 wie folgt ab:\nSetzen Sie die Ausgabe der vier logischen Funktionen Ch, Maj, Sigma0 und Sigma1 auf 0, das heißt, die\nWirkung der Funktionen wird praktisch aufgehoben.\nReduzieren Sie die Anzahl der Runden innerhalb der Kompressionsfunktion von ursprünglich 64 auf 2\nRunden.\n\nFinden Sie dann eine Kollision für die so geschwächte Version des SHA-256."
	m1Hash = Hash8(m1)
	m2Hash = Hash8(m2)
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

func Hash8(s string) string {
	result := ""

	input := []byte(s)
	hash := core.SHA256(input, 8, false)

	for _, b := range hash {
		result += fmt.Sprintf("%02x", b)
	}

	return result
}

func Hash16(s string) string {
	result := ""

	input := []byte(s)
	hash := core.SHA256(input, 16, false)

	for _, b := range hash {
		result += fmt.Sprintf("%02x", b)
	}

	return result
}
