package utils

func BytesToBits(b []byte) string {
	bits := ""
	for idx, byteVal := range b {
		for i := 7; i >= 0; i-- {
			if byteVal&(1<<i) != 0 {
				bits += "1"
			} else {
				bits += "0"
			}
		}

		// add a newline every 4 bytes
		if (idx+1)%8 == 0 {
			bits += "\n"
		}

		// add a space between bytes
		bits += " "

	}
	return bits
}
