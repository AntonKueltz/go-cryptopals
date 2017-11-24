package set1

import "encoding/hex"
import "fmt"
import "log"

func score(candidate []byte) int {
    totalScore := 0
    frequentLetters := "etaoinhs ";

    for i := 0; i < len(candidate); i++ {
        for j := 0; j < len(frequentLetters); j++ {
            if candidate[i] == frequentLetters[j] {
                totalScore++
                break
            }
        }
    }

    return totalScore
}

func singleXor(key byte, encoded []byte) []byte {
    decoded := make([]byte, len(encoded))

    for i := 0; i < len(encoded); i++ {
        decoded[i] = encoded[i] ^ key
    }

    return decoded
}

// BreakEncodedText finds the best guess plaintext for single byte key
func BreakEncodedText(encoded []byte) ([]byte, byte, int) {
    topScore, bestKey, bestGuess := 0, byte(0), []byte(nil)


    for key := 0; key < 0x100; key++ {
        decoded := singleXor(byte(key), encoded)
        keyScore := score(decoded)

        if keyScore > topScore {
            bestGuess = decoded
            bestKey = byte(key)
            topScore = keyScore
        }
    }

    return bestGuess, bestKey, topScore
}

// Challenge3Main solves set 1 challenge 3
func Challenge3Main() {
    hexString := []byte("1b37373331363f78151b7f2b783431333d78397828372d363c7" +
                        "8373e783a393b3736")

    rawBytes := make([]byte, hex.DecodedLen(len(hexString)))
    _, err := hex.Decode(rawBytes, hexString)
    if err != nil { log.Fatal(err) }

    bestGuess, _, _ := BreakEncodedText(rawBytes)
    fmt.Printf("%s\n", bestGuess)
}
