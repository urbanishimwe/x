package cryptopal

import (
	"os"
	"testing"
)

var (
	exercise1Data  = []byte("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	exercise2Data1 = []byte("1c0111001f010100061a024b53535009181c")
	exercise2Data2 = []byte("686974207468652062756c6c277320657965")
	exercise3Data  = []byte("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	exercise5Data  = []byte("Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal")
	exercise5Key   = []byte("ICE")
	exercise7key   = []byte("YELLOW SUBMARINE")
)

const (
	exercise1Res = "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	exercise2Res = "746865206b696420646f6e277420706c6179"
	exercise3Res = "Cooking MC's like a pound of bacon"
	exercise4Res = "Now that the party is jumping\n"
	exercise5Res = `0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f`
	exercise8Res = "d880619740a8a19b7840a8a31c810a3d08649af70dc06f4fd5d2d69c744cd283e2dd052f6b641dbf9d11b0348542bb5708649af70dc06f4fd5d2d69c744cd2839475c9dfdbc1d46597949d9c7e82bf5a08649af70dc06f4fd5d2d69c744cd28397a93eab8d6aecd566489154789a6b0308649af70dc06f4fd5d2d69c744cd283d403180c98c8f6db1f2a3f9c4040deb0ab51b29933f2c123c58386b06fba186a"
)

func TestSet1(t *testing.T) {
	assert(t, exercise1(exercise1Data), exercise1Res, "exercise 1")
	assert(t, exercise2(exercise2Data1, exercise2Data2), exercise2Res, "exercise 2")
	assert(t, exercise3(exercise3Data), exercise3Res, "exercise 3")
	assert(t, exercise4(readFile(t, "testdata/set1exercise4.txt")), exercise4Res, "exercise 4")
	assert(t, exercise5(exercise5Data, exercise5Key), exercise5Res, "exercise 5")
	assert(t, exercise6(readFile(t, "testdata/set1exercise6.txt")), string(readFile(t, "testdata/set1exercise6Res.txt")), "exercise 6")
	assert(t, exercise7(readFile(t, "testdata/set1exercise7.txt"), exercise7key), string(readFile(t, "testdata/set1exercise6Res.txt"))+"\x04\x04\x04\x04", "exercise 7")
	assert(t, exercise8(readFile(t, "testdata/set1exercise8.txt")), exercise8Res, "exercise 8")
}

func assert(t *testing.T, a, b string, name string) {
	if a != b {
		t.Logf("%s Failed!", name)
		t.Fail()
	}
}

func readFile(t *testing.T, name string) []byte {
	buf, err := os.ReadFile(name)
	if err != nil {
		t.Logf("error reading %s %s", name, err)
		return nil
	}
	return buf
}
