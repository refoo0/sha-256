package main

import (
	"encoding/hex"
	"fmt"

	"github.com/refoo0/sha-256/core"
)

func main() {

	// Aufgabe 2 a)
	fmt.Println("==================================================")
	fmt.Println("Aufgabe 2 a)")
	fmt.Println("==================================================")
	message := ""
	expectedHash := "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"
	messageHash := Hash(message)
	core.ColorPrintDiff(messageHash, expectedHash)
	fmt.Printf(" - %s - Gleicher Hash: %t\n", expectedHash, messageHash == expectedHash)

	message = "abc"
	expectedHash = "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad"
	messageHash = Hash(message)
	core.ColorPrintDiff(messageHash, expectedHash)
	fmt.Printf(" - %s - Gleicher Hash: %t\n", expectedHash, messageHash == expectedHash)

	message = "abcdbcdecdefdefgefghfghighijhijkijkljklmklmnlmnomnopnopq"
	expectedHash = "248d6a61d20638b8e5c026930c3e6039a33ce45964ff2167f6ecedd419db06c1"
	messageHash = Hash(message)
	core.ColorPrintDiff(messageHash, expectedHash)
	fmt.Printf(" - %s - Gleicher Hash: %t\n", expectedHash, messageHash == expectedHash)

	// Aufgabe 2 b)
	fmt.Println("==================================================")
	fmt.Println("Aufgabe 2 b)")
	fmt.Println("==================================================")
	//    |XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX|                               |XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX|                               |XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX|                               |XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX|                               |XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX|                               |XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX|                               |XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX|                               |
	m1 := "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	m2 := "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	m1Hash := Hash8(m1)
	m2Hash := Hash8(m2)
	fmt.Printf("%s\n%s\n", m1, m2)
	core.ColorPrintDiff(m1Hash, m2Hash)
	fmt.Printf(" - %s - Gleicher Hash: %t\n", m2Hash, m1Hash == m2Hash)

	m1 = "Aufgabe 2 (Implementierung und Analyse des SHA-256)"
	m2 = "Aufgabe 2 (Implementierung und Analyse des SHA-224)"
	m1Hash = Hash8(m1)
	m2Hash = Hash8(m2)
	fmt.Printf("%s\n%s\n", m1, m2)
	core.ColorPrintDiff(m1Hash, m2Hash)
	fmt.Printf(" - %s - Gleicher Hash: %t\n", m2Hash, m1Hash == m2Hash)

	m1 = "Aufgabe 2 (Implementierung und Analyse des SHA-256)"
	m2 = "Aufgabe 3 (Implementierung und Analyse des SHA-256)"
	m1Hash = Hash8(m1)
	m2Hash = Hash8(m2)
	fmt.Printf("%s\n%s\n", m1, m2)
	core.ColorPrintDiff(m1Hash, m2Hash)
	fmt.Printf(" - %s - Gleicher Hash: %t\n", m2Hash, m1Hash == m2Hash)

	m1 = "Implementieren Sie in einer Programmiersprache Ihrer Wahl die kryptographische Hash-Funktion\nSHA-256. Falls Sie sich für die Programmiersprache C entscheiden, können Sie die zur Verfügung\ngestellte teilweise Implementierung ergänzen."
	m2 = "Implementieren Sie in einer Programmiersprache Ihrer Wahl die kryptographische Hash-Funktion\nSHA-224. Falls Sie sich für die Programmiersprache C entscheiden, können Sie die zur Verfügung\ngestellte teilweise Implementierung ergänzen."
	m1Hash = Hash8(m1)
	m2Hash = Hash8(m2)
	fmt.Printf("%s\n%s\n", m1, m2)
	core.ColorPrintDiff(m1Hash, m2Hash)
	fmt.Printf(" - %s - Gleicher Hash: %t\n", m2Hash, m1Hash == m2Hash)

	m1 = "Schwächen Sie nun Ihre Implementierung des SHA-256 wie folgt ab:\nSetzen Sie die Ausgabe der vier logischen Funktionen Ch, Maj, Sigma0 und Sigma1 auf 0, das heißt, die\nWirkung der Funktionen wird praktisch aufgehoben.\nReduzieren Sie die Anzahl der Runden innerhalb der Kompressionsfunktion von ursprünglich 64 auf 8\nRunden.\nFinden Sie dann eine Kollision für die so geschwächte Version des SHA-256."
	m2 = "Schwächen Sie nun Ihre Implementierung des SHA-224 wie folgt ab:\nSetzen Sie die Ausgabe der vier logischen Funktionen Ch, Maj, Sigma0 und Sigma1 auf 0, das heißt, die\nWirkung der Funktionen wird praktisch aufgehoben.\nReduzieren Sie die Anzahl der Runden innerhalb der Kompressionsfunktion von ursprünglich 64 auf 2\nRunden.\nFinden Sie dann eine Kollision für die so geschwächte Version des SHA-256."
	m1Hash = Hash8(m1)
	m2Hash = Hash8(m2)
	fmt.Printf("%s\n%s\n", m1, m2)
	core.ColorPrintDiff(m1Hash, m2Hash)
	fmt.Printf(" - %s - Gleicher Hash: %t\n", m2Hash, m1Hash == m2Hash)

	m1 = "Schwächen Sie nun Ihre Implementierung des SHA-256 wie folgt ab:\nSetzen Sie die Ausgabe der vier logischen Funktionen Ch, Maj, Sigma0 und Sigma1 auf 0, das heißt, die\nWirkung der Funktionen wird praktisch aufgehoben.\nReduzieren Sie die Anzahl der Runden innerhalb der Kompressionsfunktion von ursprünglich 64 auf 8\nRunden.\nFinden Sie dann eine Kollision für die so geschwächte Version des SHA-256."
	m2 = "Schwächen Sie nun Ihre Implementierung des SHA-224 wie folgt ab:\nSetzen Sie die Ausgabe der vier logischen Funktionen Ch, Maj, Sigma0 und Sigma1 auf 0, das heißt, die\nWirkung der Funktionen wird praktisch aufgehoben.\nReduzieren Sie die Anzahl der Runden innerhalb der Kompressionsfunktion von ursprünglich 64 auf 2\nRunden.\nFinden Sie dann eine Kollision für die so geschwächte Version des SHA-224."
	m1Hash = Hash8(m1)
	m2Hash = Hash8(m2)
	fmt.Printf("%s\n%s\n", m1, m2)
	core.ColorPrintDiff(m1Hash, m2Hash)
	fmt.Printf(" - %s - Gleicher Hash: %t\n", m2Hash, m1Hash == m2Hash)

	m1 = "Schwächen Sie nun Ihre Implementierung des SHA-256 wie folgt ab:\nSetzen Sie die Ausgabe der vier logischen Funktionen Ch, Maj, Sigma0 und Sigma1 auf 0, das heißt, die\nWirkung der Funktionen wird praktisch aufgehoben.\nReduzieren Sie die Anzahl der Runden innerhalb der Kompressionsfunktion von ursprünglich 64 auf 8\nRunden.\nFinden Sie dann eine Kollision für die so geschwächte Version des SHA-256."
	m2 = "Schwächen Sie nun Ihre Implementierung des SHA-256 wie folgt ab:\nSetzen Sie die Ausgabe der vier logischen Funktionen Ch, Maj, Sigma0 und Sigma1 auf 0, das heißt, die\nWirkung der Funktionen wird praktisch aufgehoben.\nReduzieren Sie die Anzahl der Runden innerhalb der Kompressionsfunktion von ursprünglich 64 auf 8\nRunden.\nFinden Sie dann eine Kollision für die so geschwächte Version des SHA-256.abc"
	m1Hash = Hash8(m1)
	m2Hash = Hash8(m2)
	fmt.Printf("%s\n%s\n", m1, m2)
	core.ColorPrintDiff(m1Hash, m2Hash)
	fmt.Printf(" - %s - Gleicher Hash: %t\n", m2Hash, m1Hash == m2Hash)

	// Aufgabe 2 c)
	fmt.Println("==================================================")
	fmt.Println("Aufgabe 2 c)")
	fmt.Println("==================================================")
	b1 := []byte{
		0x00, 0x00, 0x00, 0x00, // W0
		0x00, 0x00, 0x00, 0x52,
		0x00, 0x00, 0x00, 0x5b,
		0x00, 0x00, 0x00, 0x00,

		0x00, 0x00, 0x00, 0x40, // W4
		0x00, 0x00, 0x00, 0x02,
		0x00, 0x00, 0x00, 0x05,
		0x00, 0x00, 0x00, 0x60,

		0x00, 0x00, 0x00, 0xa2, // W8
		0x00, 0x00, 0x00, 0xe0,
		0x00, 0x00, 0x00, 0x05,
		0x00, 0x00, 0x00, 0x05,

		0x00, 0x00, 0x00, 0x64, // W12
		0x00, 0x00, 0x00, 0x80,
		0x00, 0x00, 0x00, 0x54,
		0x00, 0x00, 0x00, 0x08,
	}

	b2 := []byte{
		0x00, 0x00, 0x00, 0x00, // W0
		0x00, 0x00, 0x00, 0x52,
		0x00, 0x00, 0x00, 0x5b,
		0x00, 0x00, 0x00, 0x00,

		0x00, 0x00, 0x00, 0x40, // W4
		0x00, 0x00, 0x00, 0x82,
		0x00, 0x00, 0x00, 0x05,
		0x00, 0x00, 0x00, 0x60,

		0x00, 0x00, 0x00, 0xa2, // W8
		0x00, 0x00, 0x00, 0x60,
		0x00, 0x00, 0x00, 0x05,
		0x00, 0x00, 0x00, 0x05,

		0x00, 0x00, 0x00, 0x64, // W12
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x54,
		0x00, 0x00, 0x00, 0x08,
	}
	hb1 := core.SHA256(b1, 16, false, false)
	hb2 := core.SHA256(b2, 16, false, false)
	fmt.Print("Nachricht 1: ")
	core.ColorPrintDiff(hex.EncodeToString(b1), hex.EncodeToString(b2))
	fmt.Print("\nNachricht 2: ")
	fmt.Printf("%s\n", hex.EncodeToString(b2))
	for i := range len(hb1) {
		if hb1[i] == hb2[i] {
			core.ColorPrintf(core.Green, "%02x", hb1[i])
		} else {
			core.ColorPrintf(core.Red, "%02x", hb1[i])
		}
	}
	fmt.Printf("\n%02x - Gleicher Hash: %t\n", hb1, hb1 == hb2)

	b1 = []byte{
		0x27, 0xf5, 0xd0, 0xfa,
		0x00, 0x01, 0x7f, 0xc4,
		0x00, 0x00, 0x00, 0x01,
		0x00, 0x00, 0x00, 0x01,

		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,

		0xf0, 0xf5, 0xfe, 0xf9,
		0xff, 0xfc, 0xff, 0xfb,
		0xff, 0xff, 0xff, 0xfc,
		0xff, 0xff, 0xff, 0xfd,

		0x5f, 0xf0, 0xc0, 0x00,
		0xff, 0xff, 0xff, 0xfb,
		0x00, 0x00, 0x00, 0x03,
		0xff, 0xff, 0xff, 0xff,
	}

	b2 = []byte{
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,

		0x80, 0xe1, 0x14, 0xf9,
		0xff, 0xff, 0xff, 0xbf,
		0x00, 0x00, 0x00, 0x03,
		0x00, 0x00, 0x00, 0x00,

		0xc0, 0x00, 0x8b, 0xf4,
		0xff, 0xff, 0xff, 0xc4,
		0xff, 0xff, 0xff, 0xfb,
		0xff, 0xff, 0xff, 0xff,

		0x07, 0x05, 0x7c, 0x01,
		0x00, 0x01, 0x80, 0x00,
		0x00, 0x00, 0x00, 0x01,
		0x00, 0x00, 0x00, 0x00,
	}
	hb1 = core.SHA256(b1, 16, false, false)
	hb2 = core.SHA256(b2, 16, false, false)
	fmt.Print("Nachricht 1: ")
	core.ColorPrintDiff(hex.EncodeToString(b1), hex.EncodeToString(b2))
	fmt.Print("\nNachricht 2: ")
	fmt.Printf("%s\n", hex.EncodeToString(b2))
	for i := range len(hb1) {
		if hb1[i] == hb2[i] {
			core.ColorPrintf(core.Green, "%02x", hb1[i])
		} else {
			core.ColorPrintf(core.Red, "%02x", hb1[i])
		}
	}
	fmt.Printf("\n%02x - Gleicher Hash: %t\n", hb1, hb1 == hb2)

	m1 = "IXBZ=AD7w.Ay?]\"D^xZ A; n!N31{ !G/iBKH_*0.0xo<mS QRRaki~ ]-}8;QyO"
	m2 = "\"}[X=8/AR>hyBp!BuXDOVkK y8.1=45 f?& 3A)j &/ot3AKaMO4VBhd*3[8v*fx"
	m1Hash = Hash16(m1)
	m2Hash = Hash16(m2)
	fmt.Print("Nachricht 1: ")
	core.ColorPrintDiff(m1, m2)
	fmt.Print("\nNachricht 2: ")
	fmt.Printf("%s\n", m2)
	// fmt.Printf("%s\n%s\n", m1, m2)
	core.ColorPrintDiff(m1Hash, m2Hash)
	fmt.Printf("\n%s - Gleicher Hash: %t\n", m2Hash, m1Hash == m2Hash)

	m1 = "PYL+Vpg~|}~{~ ?P\"WjzR!s~_5M*!j\\Rpq\\r^gsO@3~wP.pPXc<`2]toUU( dPH!"
	m2 = "[r@2yRg)ve$&} v}Oabo8|w _XX!Gx4&-5|o2HpWLA(*, *\" @Pd) q\"[Jw~?B9 "
	m1Hash = Hash16(m1)
	m2Hash = Hash16(m2)
	fmt.Print("Nachricht 1: ")
	core.ColorPrintDiff(m1, m2)
	fmt.Print("\nNachricht 2: ")
	fmt.Printf("%s\n", m2)
	// fmt.Printf("%s\n%s\n", m1, m2)
	core.ColorPrintDiff(m1Hash, m2Hash)
	fmt.Printf("\n%s - Gleicher Hash: %t\n", m2Hash, m1Hash == m2Hash)
}

func Hash(s string) string {
	result := ""

	input := []byte(s)
	hash := core.SHA256(input, 64, true, true)

	for _, b := range hash {
		result += fmt.Sprintf("%02x", b)
	}

	return result
}

func Hash8(s string) string {
	result := ""

	input := []byte(s)
	hash := core.SHA256(input, 8, false, false)

	for _, b := range hash {
		result += fmt.Sprintf("%02x", b)
	}

	return result
}

func Hash16(s string) string {
	result := ""

	input := []byte(s)
	hash := core.SHA256(input, 16, false, false)

	for _, b := range hash {
		result += fmt.Sprintf("%02x", b)
	}

	return result
}
