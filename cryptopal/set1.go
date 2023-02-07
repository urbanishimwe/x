package cryptopal

import (
	"bufio"
	"bytes"
	"crypto/aes"
	"encoding/base64"
	"encoding/hex"
	"math"
	"math/bits"
)

// Convert hex to base64
func exercise1(encoded []byte) string {
	n, _ := hex.Decode(encoded, encoded)
	return base64.RawStdEncoding.EncodeToString(encoded[:n])
}

// Fixed XOR
func exercise2(a, b []byte) string {
	hex.Decode(a, a)
	n, _ := hex.Decode(b, b)
	for i := 0; i < n; i++ {
		a[i] ^= b[i]
	}
	return hex.EncodeToString(a[:n])
}

// crack single-byte XOR cipher (it reduces the size of the output)
func exercise3(buf []byte) string {
	n, _ := hex.Decode(buf, buf)
	var tmp = make([]byte, n)
	var res = make([]byte, n)
	var prevScore int
	for x := 0; x < 256; x++ {
		for i := 0; i < n; i++ {
			tmp[i] = buf[i] ^ byte(x)
		}
		score := asciiPrintablity(tmp)
		if score > prevScore {
			copy(res, tmp)
			prevScore = score
		}
	}
	// 88 "Cooking MC's like a pound of bacon"
	return string(res)
}

// Detect single-character XOR from file lines
func exercise4(buf []byte) string {
	rd := bufio.NewReaderSize(bytes.NewReader(buf), len(buf))
	score := 0
	res := ""
	for {
		s, _ := rd.ReadBytes('\n')
		if len(s) == 0 {
			break
		}
		if s[len(s)-1] == '\n' {
			s = s[:len(s)-1]
		}
		n := exercise3(s)
		if p := asciiPrintablity([]byte(n)); p > score {
			score = p
			res = n
		}
	}
	return res
}

// implement repeating-key XOR
func exercise5(word, key []byte) string {
	n := len(key) // key is "ICE"
	for i := 0; i < len(word); i++ {
		word[i] ^= key[i%n]
	}
	return hex.EncodeToString(word)
}

// crack repeating-key XOR
func exercise6(c []byte) string {
	c = b64toS(string(c))
	keySize := keySizeGuess(c, 2, 40)
	// guessing key
	key := make([]byte, keySize)
	for i := 0; i < keySize; i++ {
		var bestK, bestMatch int
		for k := 0; k < 256; k++ {
			match := 0
			for ic := i; ic < len(c); ic += keySize {
				xor := c[ic] ^ byte(k)
				if dict[xor] == 1 {
					match++
				}
			}
			if match > bestMatch {
				bestK = k
				bestMatch = match
			}
		}
		key[i] = byte(bestK)
		bestMatch = 0
	}

	// decypher
	dec := make([]byte, len(c))
	for i := 0; i < len(c); i++ {
		dec[i] = c[i] ^ key[i%keySize]
	}
	return string(dec)
}

// decrypt AES-128-ECB
func exercise7(data, key []byte) string {
	block, _ := aes.NewCipher(key)
	text := b64toS(string(data))
	var line [aes.BlockSize]byte
	var res []byte
	for i := 0; i < len(text); i += aes.BlockSize {
		block.Decrypt(line[:], text[i:i+aes.BlockSize])
		res = append(res, line[:]...)
	}
	return string(res)
}

func exercise8(data []byte) string {
	var best []byte
	bestTimes := 0
	rd := bufio.NewReaderSize(bytes.NewReader(data), len(data))
	for {
		b, _ := rd.ReadBytes('\n')
		if len(b) == 0 {
			break
		}
		if b[len(b)-1] == '\n' {
			b = b[:len(b)-1]
		}
		var tmp [aes.BlockSize]byte
		tmpTotal := 0
		m := make(map[[aes.BlockSize]byte]int)
		for i := 0; i < len(b); i += aes.BlockSize {
			copy(tmp[:], b[i:i+aes.BlockSize])
			m[tmp]++
			if m[tmp] > 1 {
				tmpTotal++
			}
		}
		if tmpTotal > bestTimes {
			best = b
			bestTimes = tmpTotal
		}
	}
	return string(best)
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

func hammingDistance(a, b []byte) (dist int) {
	for i := 0; i < len(a); i++ {
		dist += bits.OnesCount8(a[i] ^ b[i])
	}
	return
}

func b64toS(b64 string) []byte {
	s, _ := base64.StdEncoding.DecodeString(b64)
	return s
}

func keySizeGuess(c []byte, a, b int) int {
	key := math.MaxInt
	editDist := math.MaxInt
	for i := a; i <= b; i++ {
		sum := 0
		// test on 10 blocks
		for j := 0; j < 10; j++ {
			sum += hammingDistance(c[i*j:j*i+i], c[j*i+i:j*i+i*2])
		}
		sum /= i
		if sum < editDist {
			editDist = sum
			key = i
		}
	}
	return key
}
