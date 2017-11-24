package set1

import "encoding/hex"
import "errors"
import "fmt"
import "log"

// Xor performs logical XOR on two byte arrays
func Xor(bytes1, bytes2 []byte) ([]byte, error) {
    if len(bytes1) != len(bytes2) {
        return []byte(nil), errors.New("Byte arrays must bt the same length")
    }

    byteLen := len(bytes1)
    result := make([]byte, byteLen)

    for i := 0; i < byteLen; i++ {
        result[i] = bytes1[i] ^ bytes2[i]
    }

    return result, nil
}


// Challenge2Main solves set 1 challenge 2
func Challenge2Main() {
    hexString1 := []byte("1c0111001f010100061a024b53535009181c")
    hexString2 := []byte("686974207468652062756c6c277320657965")

    rawBytes1 := make([]byte, hex.DecodedLen(len(hexString1)))
    _, err := hex.Decode(rawBytes1, hexString1)
    if err != nil { log.Fatal(err) }

    rawBytes2 := make([]byte, hex.DecodedLen(len(hexString2)))
    _, err = hex.Decode(rawBytes2, hexString2)
    if err != nil { log.Fatal(err) }

    result, _ := Xor(rawBytes1, rawBytes2)

    hexEncoded := make([]byte, hex.EncodedLen(len(result)))
    hex.Encode(hexEncoded, result)
    fmt.Printf("%s\n", hexEncoded)
}
