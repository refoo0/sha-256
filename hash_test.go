package main

import "testing"

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

}
