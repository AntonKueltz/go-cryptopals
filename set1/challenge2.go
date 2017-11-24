package set1

import "encoding/hex"
import "fmt"
import "log"

func min(a, b int) int {
    if a <= b {
        return a
    }
    return b
}

// Challenge2Main solves set 1 challenge 2
func Challenge2Main() {
    hexString1 := []byte("1c0111001f010100061a024b53535009181c")
    hexString2 := []byte("686974207468652062756c6c277320657965")

    rawBytes1 := make([]byte, hex.DecodedLen(len(hexString1)))
    rawLen1, err := hex.Decode(rawBytes1, hexString1)
    if err != nil { log.Fatal(err) }

    rawBytes2 := make([]byte, hex.DecodedLen(len(hexString2)))
    rawLen2, err := hex.Decode(rawBytes2, hexString2)
    if err != nil { log.Fatal(err) }

    smallerLen := min(rawLen1, rawLen2)
    result := make([]byte, smallerLen)
    for i := 0; i < smallerLen; i++ {
        result[i] = rawBytes1[i] ^ rawBytes2[i]
    }

    hexEncoded := make([]byte, hex.EncodedLen(len(result)))
    hex.Encode(hexEncoded, result)
    fmt.Printf("%s\n", hexEncoded)
}
