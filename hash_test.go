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

func TestHashCollision(t *testing.T) {
	input1 := "ab"
	input2 := "ba"
	hash1 := core.SHA256([]byte(input1), 8, false)
	hash2 := core.SHA256([]byte(input2), 8, false)

	if !bytes.Equal(hash1[:], hash2[:]) {
		t.Errorf("Hash collision detected: %x != %x", hash1, hash2)
	}

}
