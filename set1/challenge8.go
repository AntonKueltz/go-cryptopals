package set1

import "bytes"
import "crypto/aes"
import "encoding/hex"
import "fmt"
import "io/ioutil"
import "log"
import "path/filepath"

// DetectEcbMode indicates if a ciphertext was encrypted in ECB mode
func DetectEcbMode(ciphertext []byte) bool {
    blockCount := len(ciphertext) / aes.BlockSize
    var blocks [][]byte

    for i := 0; i < blockCount; i++ {
        start, end := i * aes.BlockSize, (i + 1) * aes.BlockSize
        block := ciphertext[start:end]

        for _, otherBlock := range blocks {
            if bytes.Compare(block, otherBlock) == 0 {
                return true
            }
        }

        blocks = append(blocks, block)
    }

    return false
}

// Challenge8Main solves set 1 challenge 8
func Challenge8Main() {
    fileLocation, err := filepath.Abs("set1/files/8.txt")
    if err != nil { log.Fatal(err) }

    fileBytes, err := ioutil.ReadFile(fileLocation)
    if err != nil { log.Fatal(err) }

    splitLines := bytes.Split(fileBytes, []byte("\n"))
    for _, line := range splitLines {
        rawBytes := make([]byte, hex.DecodedLen(len(line)))
        _, err := hex.Decode(rawBytes, line)
        if err != nil { log.Fatal(err) }

        if DetectEcbMode(rawBytes) {
            fmt.Printf("%s\n", line)
            break
        }
    }
}
