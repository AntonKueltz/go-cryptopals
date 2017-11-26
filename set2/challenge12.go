package set2

import "encoding/base64"
import "fmt"
import "strings"

import "github.com/AntonKueltz/go-cryptopals/set1"
// EncryptAesEcb, DetectEcbMode

var c12Key = []byte("COUNTERATTACKERS")  // or something else random...
var b64 = "Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkg" +
          "aGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBq" +
          "dXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUg" +
          "YnkK"
var target, _ = base64.StdEncoding.DecodeString(b64)

func encrypt(prefix []byte) []byte {
    plaintext := make([]byte, len(prefix) + len(target))
    copy(plaintext[:len(prefix)], prefix)
    copy(plaintext[len(prefix):], target)

    ciphertext := set1.EncryptAesEcb(c12Key, plaintext)
    return ciphertext
}

func discoverBlockSize() int {
    prefix := []byte(nil)
    initialLen := len(encrypt(prefix))
    newLen := initialLen

    for newLen == initialLen {
        prefix = append(prefix, byte('A'))
        newLen = len(encrypt(prefix))
    }

    return newLen - initialLen
}

func buildLookupTable(prefix string, blocksize int) map[string]byte {
    lookup := make(map[string]byte)

    for i := 0; i < 0x100; i++ {
        block := append([]byte(prefix), byte(i))
        ciphertext := encrypt(block)
        lookup[string(ciphertext[:blocksize])] = byte(i)
    }

    return lookup
}

func breakByteByByte(blocksize int) []byte {
    prefix := strings.Repeat("A", 16)
    known := ""
    totalBytes := len(encrypt([]byte(nil)))
    plaintext := make([]byte, totalBytes)


    for i := 0; i < totalBytes; i++ {
        if i > 0 && i % blocksize == 0 {
            known = ""
            prefix = string(plaintext[(i - blocksize):i])
        }

        prefixBytes := (blocksize - 1) - (i % blocksize)
        shiftedPrefix := prefix[(blocksize - prefixBytes):]
        lookup := buildLookupTable(shiftedPrefix + known, blocksize)
        ciphertext := encrypt([]byte(shiftedPrefix))

        start := (i / blocksize) * blocksize
        lookupBytes := ciphertext[start:(start + blocksize)]
        plaintextByte := lookup[string(lookupBytes)]
        plaintext[i] = plaintextByte

        known += string(plaintextByte)
    }

    return plaintext
}

// Challenge12Main solves set 2 challenge 12
func Challenge12Main() {
    blocksize := discoverBlockSize()
    fmt.Printf("Detected blocksize of %d...\n", blocksize)

    repeating := encrypt([]byte(strings.Repeat("A", 100)))
    if set1.DetectEcbMode(repeating) {
        fmt.Printf("Detected ECB mode...\n");
    }

    plaintext := breakByteByByte(blocksize)
    fmt.Printf("%s\n", plaintext)
}
