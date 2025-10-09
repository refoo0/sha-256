package core

import (
	"fmt"
)

func expandMessage(block []byte, verbose bool) [64]uint32 {

	if verbose {
		fmt.Println("--- Expanding the Message into 64 words ---")
		fmt.Println("First 16 words are directly from the block, the remaining 48 words are generated from the first 16 words")

		fmt.Println("    s0 := (w[i-15] right rotate 7) xor (w[i-15] right rotate 18) xor (w[i-15] right shift 3)")
		fmt.Println("    s1 := (w[i-2] right rotate 17) xor (w[i-2] right rotate 19) xor (w[i-2] right shift 10)")
		fmt.Println("    w[i] := w[i-16] + s0 + w[i-7] + s1")
	}

	var w [64]uint32
	// First 16 words are directly from the block
	for i := 0; i < 16; i++ {
		w[i] = uint32(block[i*4]) << 24
		w[i] += uint32(block[i*4+1]) << 16
		w[i] += uint32(block[i*4+2]) << 8
		w[i] += uint32(block[i*4+3])

		if verbose {
			fmt.Println("W[" + fmt.Sprintf("%02d", i) + "] = " + fmt.Sprintf("%032b", w[i]))
		}
	}

	// Extend the first 16 words into the remaining 48 words w[16..63] of the message schedule array
	// for i from 16 to 63 do
	//     s0 := (w[i-15] rightrotate 7) xor (w[i-15] rightrotate 18) xor (w[i-15] rightshift 3)
	//     s1 := (w[i-2] rightrotate 17) xor (w[i-2] rightrotate 19) xor (w[i-2] rightshift 10)
	//     w[i] := w[i-16] + s0 + w[i-7] + s1
	for i := 16; i < 64; i++ {
		s0 := Sigma0(w[i-15])
		s1 := Sigma1(w[i-2])
		w[i] = w[i-16] + s0 + w[i-7] + s1
		if verbose {
			fmt.Println("W[" + fmt.Sprintf("%02d", i) + "] = " + fmt.Sprintf("%032b", w[i]))
		}
	}
	return w
}

func compress(hashes [8]uint32, w [64]uint32, rounding bool, iterations int, verbose bool) [8]uint32 {
	if verbose {
		fmt.Println("--- Rounding through 64 iterations ---")
		fmt.Println("For each of the 64 rounds:")
		fmt.Println("    for first block there are 8 initial hash values ")
		for i := 0; i < 8; i++ {
			fmt.Println("        H[" + fmt.Sprintf("%d", i) + "] = " + fmt.Sprintf("%032b", hashes[i]) + " " + fmt.Sprintf("%08x", hashes[i]))
		}

		fmt.Println("    S1 := (w[i] right rotate 6) xor (w[i] right rotate 11) xor (w[i] right rotate 25)")
		fmt.Println("    ch := (w[i] and f) xor ((not w[i]) and g)")
		fmt.Println("    temp1 := h + S1 + ch + k[i] + w[i]")
		fmt.Println("    S0 := (a right rotate 2) xor (a right rotate 13) xor (a right rotate 22)")
		fmt.Println("    maj := (a and b) xor (a and c) xor (b and c)")
		fmt.Println("    temp2 := S0 + maj")
	}

	a, b, c, d, e, f, g, h := hashes[0], hashes[1], hashes[2], hashes[3], hashes[4], hashes[5], hashes[6], hashes[7]

	if verbose {
		fmt.Println()
		fmt.Println("    h := g")
		fmt.Println("    g := f")
		fmt.Println("    f := e")
		fmt.Println("    e := d + temp1")
		fmt.Println("    d := c")
		fmt.Println("    c := b")
		fmt.Println("    b := a")
		fmt.Println("    a := temp1 + temp2")
	}

	// Main loop
	for i := 0; i < iterations; i++ {
		s1 := BigSigma1(e)
		ch := Ch(e, f, g)
		var tmp uint32
		if rounding {
			tmp = h + s1 + ch + Sha256Constants[i] + w[i]
		} else {
			tmp = h + Sha256Constants[i] + w[i]
		}

		s0 := BigSigma0(a)
		maj := Maj(a, b, c)

		var tmp2 uint32
		if rounding {
			tmp2 = tmp + s0 + maj
		} else {
			tmp2 = tmp
		}

		h = g
		g = f
		f = e
		e = d + tmp
		d = c
		c = b
		b = a
		a = tmp2

		/* 		fmt.Println("Round " + fmt.Sprintf("%02d", i+1) + ":")
		   		fmt.Println("    a = " + fmt.Sprintf("%032b", a) + " " + fmt.Sprintf("%08x", a))
		   		fmt.Println("    b = " + fmt.Sprintf("%032b", b) + " " + fmt.Sprintf("%08x", b))
		   		fmt.Println("    c = " + fmt.Sprintf("%032b", c) + " " + fmt.Sprintf("%08x", c))
		   		fmt.Println("    d = " + fmt.Sprintf("%032b", d) + " " + fmt.Sprintf("%08x", d))
		   		fmt.Println("    e = " + fmt.Sprintf("%032b", e) + " " + fmt.Sprintf("%08x", e))
		   		fmt.Println("    f = " + fmt.Sprintf("%032b", f) + " " + fmt.Sprintf("%08x", f))
		   		fmt.Println("    g = " + fmt.Sprintf("%032b", g) + " " + fmt.Sprintf("%08x", g))
		   		fmt.Println("    h = " + fmt.Sprintf("%032b", h) + " " + fmt.Sprintf("%08x", h)) */

	}

	hashes[0] += a
	hashes[1] += b
	hashes[2] += c
	hashes[3] += d
	hashes[4] += e
	hashes[5] += f
	hashes[6] += g
	hashes[7] += h

	if verbose {
		fmt.Println("")
		for i := 0; i < 8; i++ {
			fmt.Println("H[" + fmt.Sprintf("%d", i) + "] = " + fmt.Sprintf("%032b", hashes[i]) + " " + fmt.Sprintf("%08x", hashes[i]))
		}
	}
	return hashes
}

func SHA256(input []byte, iterations int, rounding bool, verbose bool) [32]byte {
	// IV-Values
	hashes := InitialHashes

	paddedMessage := PadMessage(input, verbose)

	for n := 0; n < len(paddedMessage); n += 64 {
		block := paddedMessage[n : n+64]
		if verbose {
			fmt.Println("--- Separating the Message into 512-bit blocks ---")
			fmt.Println("Processing block " + fmt.Sprintf("%d", n/64+1) + " of " + fmt.Sprintf("%d", len(paddedMessage)/64) + ":\n" + BytesToBits(block))
		}

		w := expandMessage(block, verbose)
		hashes = compress(hashes, w, rounding, iterations, verbose)

	}

	return transformToByteArray(hashes)
}
