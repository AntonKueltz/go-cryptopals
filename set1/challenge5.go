package set1

import "encoding/hex"
import "fmt"

// RepeatingKeyXor returns some text xored against a repeating key
func RepeatingKeyXor(key, text []byte) []byte {
    xored := make([]byte, len(text))

    for i := 0; i < len(text); i++ {
        xored[i] = text[i] ^ key[i % len(key)]
    }

    return xored
}

// Challenge5Main solves set 1 challenge 5
func Challenge5Main() {
    key := []byte("ICE")
    text := []byte("Burning 'em, if you ain't quick and nimble\n" +
                   "I go crazy when I hear a cymbal")

    xored := RepeatingKeyXor(key, text)
    hexEncoded := make([]byte, hex.EncodedLen(len(xored)))
	hex.Encode(hexEncoded, xored)

	fmt.Printf("%s\n", hexEncoded)
}
