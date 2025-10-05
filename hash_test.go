package main

import (
	"bytes"
	"crypto"
	"testing"

	"github.com/refoo0/sha-256/core"
)

func TestHash(t *testing.T) {

	inputs := []string{
		"",
		"abc",
		"abcdbcdecdefdefgefghfghighijhijkijkljklmklmnlmnomnopnopq",
	}

	outputs := []string{
		"e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
		"ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad",
		"248d6a61d20638b8e5c026930c3e6039a33ce45964ff2167f6ecedd419db06c1",
	}

	for i, input := range inputs {
		got := Hash(input)
		if got != outputs[i] {
			t.Errorf("Hash(%q) = %q; want %q", input, got, outputs[i])
		}
	}

	// check with crypto/sha256
	msg := "lorem ipsum dolor sit amet consectetur adipiscing elit lorem ipsum dolor sit amet consectetur adipiscing elit sed do eiusmod tempor incididuntlorem ipsum dolor sit amet consectetur adipiscing elit sed do eiusmod tempor incididuntlorem ipsum dolor sit amet consectetur adipiscing elit sed do eiusmod tempor incididuntlorem ipsum dolor sit amet consectetur adipiscing elit sed do eiusmod tempor incididunt lorem ipsum dolor sit amet consectetur adipiscing elit sed do eiusmod tempor incididunt lorem ipsum dolor sit amet consectetur adipiscing elit sed do eiusmod tempor incididunt lorem ipsum dolor sit amet consectetur adipiscing elit sed do eiusmod tempor incididunt ut labore et dolore magna aliqua ut labore et dolore magna aliqua ut labore et dolore magna aliqua ut labore et dolore magna aliqua  ut labore et dolore magna aliqua  ut labore et dolore magna aliqua  ut labore et dolore magna aliqua sed do eiusmod tempor incididunt ut labore et dolore magna aliqua "
	//msg := "a"
	expected := crypto.SHA256.New()
	expected.Write([]byte(msg))
	expectedSum := expected.Sum(nil)

	gotArr := core.SHA256([]byte(msg), 64, true)
	gotSum := gotArr[:]

	if !bytes.Equal(gotSum, expectedSum) {
		t.Errorf("Hash(%q) = %q; want %q", msg, gotSum, expectedSum)
	}

}

func TestHashCollision8(t *testing.T) {
	m1 := "Implementieren Sie in einer Programmiersprache Ihrer Wahl die kryptographische Hash-Funktion\nSHA-256. Falls Sie sich für die Programmiersprache C entscheiden, können Sie die zur Verfügung\ngestellte teilweise Implementierung ergänzen."
	m2 := "Implementieren Sie in einer Programmiersprache Ihrer Wahl die kryptographische Hash-Funktion\nSHA-224. Falls Sie sich für die Programmiersprache C entscheiden, können Sie die zur Verfügung\ngestellte teilweise Implementierung ergänzen."
	m1Hash := core.SHA256([]byte(m1), 8, false)
	m2Hash := core.SHA256([]byte(m2), 8, false)

	if m1Hash != m2Hash {
		t.Errorf("Hash8(%q) = %q; Hash8(%q) = %q; want equal", m1, m1Hash, m2, m2Hash)
	}

	m1 = "Schwächen Sie nun Ihre Implementierung des SHA-256 wie folgt ab:\nSetzen Sie die Ausgabe der vier logischen Funktionen Ch, Maj, Sigma0 und Sigma1 auf 0, das heißt, die\nWirkung der Funktionen wird praktisch aufgehoben.\nReduzieren Sie die Anzahl der Runden innerhalb der Kompressionsfunktion von ursprünglich 64 auf 8\nRunden.\n\nFinden Sie dann eine Kollision für die so geschwächte Version des SHA-256."
	m2 = "Schwächen Sie nun Ihre Implementierung des SHA-224 wie folgt ab:\nSetzen Sie die Ausgabe der vier logischen Funktionen Ch, Maj, Sigma0 und Sigma1 auf 0, das heißt, die\nWirkung der Funktionen wird praktisch aufgehoben.\nReduzieren Sie die Anzahl der Runden innerhalb der Kompressionsfunktion von ursprünglich 64 auf 2\nRunden.\n\nFinden Sie dann eine Kollision für die so geschwächte Version des SHA-256."
	m1Hash = core.SHA256([]byte(m1), 8, false)
	m2Hash = core.SHA256([]byte(m2), 8, false)

	if m1Hash != m2Hash {
		t.Errorf("Hash8(%q) = %q; Hash8(%q) = %q; want equal", m1, m1Hash, m2, m2Hash)
	}
}

func TestHashCollision16(t *testing.T) {
	m1 := "Implementieren Sie in einer Programmiersprache Ihrer Wahl die kryptographische Hash-Funktion\nSHA-256. Falls Sie sich für die Programmiersprache C entscheiden, können Sie die zur Verfügung\ngestellte teilweise Implementierung ergänzen."
	m2 := "Implementieren Sie in einer Programmiersprache Ihrer Wahl die kryptographische Hash-Funktion\nSHA-224. Falls Sie sich für die Programmiersprache C entscheiden, können Sie die zur Verfügung\ngestellte teilweise Implementierung ergänzen."
	m1Hash := core.SHA256([]byte(m1), 16, false)
	m2Hash := core.SHA256([]byte(m2), 16, false)

	if m1Hash != m2Hash {
		t.Errorf("Hash16(%q) = %q; Hash16(%q) = %q; want equal", m1, m1Hash, m2, m2Hash)
	}

}
