package set1

import "encoding/base64"
import "encoding/hex"
import "fmt"
import "log"

// Challenge1Main solves set 1 challenge 1
func Challenge1Main() {
    hexEncoded := []byte("49276d206b696c6c696e6720796f757220627261696e206c6" +
                          "96b65206120706f69736f6e6f7573206d757368726f6f6d")

    rawBytes := make([]byte, hex.DecodedLen(len(hexEncoded)))
    rawLen, err := hex.Decode(rawBytes, hexEncoded)
    if err != nil { log.Fatal(err) }

    base64Encoded := base64.StdEncoding.EncodeToString(rawBytes[:rawLen])
    fmt.Println(base64Encoded)
}
