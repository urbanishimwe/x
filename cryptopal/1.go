package cryptopal

import (
	"encoding/base64"
	"encoding/hex"
)

// Convert hex to base64
func _1(encoded []byte) string {
	n, _ := hex.Decode(encoded, encoded)
	return base64.RawStdEncoding.EncodeToString(encoded[:n])
}

// Fixed XOR
func _2(a, b []byte) string {
	n, _ := hex.Decode(a, a)
	n, _ = hex.Decode(b, b)
	for i := 0; i < n; i++ {
		a[i] ^= b[i]
	}
	return hex.EncodeToString(a[:n])
}

// Single-byte XOR cipher (it reduces the size of the output)
func _3(buf []byte) {
	n, _ := hex.Decode(buf, buf)
	var tmp = make([]byte, n)
	var res = make([]byte, n)
	var prevScore int
	var xor int
	for x := 0; x < 256; x++ {
		for i := 0; i < n; i++ {
			tmp[i] = buf[i] ^ byte(x)
		}
		score := asciiPrintablity(tmp)
		if score > prevScore {
			copy(res, tmp)
			prevScore = score
			xor = x
		}
	}
	// 88 Cooking MC's like a pound of bacon
	println(xor, string(res))
}

// Detect single-character XOR
func _4([]byte) {
	// its "Now that the party is jumping" with xor 53, at line 171
	// loop over _3() on each line and rate the result obtained using the same impls from _3()
}

// implement repeating-key XOR
func _5(word, key []byte) string {
	n := len(key) // key is "ICE"
	for i := 0; i < len(word); i++ {
		word[i] ^= key[i%n]
	}
	return hex.EncodeToString(word)
}

var dict = [256]byte{
	'a': 1, 'b': 1, 'c': 1, 'd': 1, 'e': 1, 'f': 1,
	'g': 1, 'h': 1, 'i': 1, 'j': 1, 'k': 1, 'l': 1,
	'm': 1, 'n': 1, 'o': 1, 'p': 1, 'q': 1, 'r': 1,
	's': 1, 't': 1, 'u': 1, 'v': 1, 'w': 1, 'x': 1,
	'y': 1, 'z': 1, 'A': 1, 'B': 1, 'C': 1, 'D': 1,
	'E': 1, 'F': 1, 'G': 1, 'H': 1, 'I': 1, 'J': 1,
	'K': 1, 'L': 1, 'M': 1, 'N': 1, 'O': 1, 'P': 1,
	'Q': 1, 'R': 1, 'S': 1, 'T': 1, 'U': 1, 'V': 1,
	'W': 1, 'X': 1, 'Y': 1, 'Z': 1, '0': 1, '1': 1,
	'2': 1, '3': 1, '4': 1, '5': 1, '6': 1, '7': 1,
	'8': 1, '9': 1, '!': 1, ' ': 1, '\'': 1, '?': 1,
	'.': 1, ',': 1,
}

func asciiPrintablity(buf []byte) (score int) {
	for i := 0; i < len(buf); i++ {
		// this is clever, yet powerful for non-chinese words
		if dict[buf[i]] > 0 {
			score++
		}
	}
	return
}
