package set1

import "bytes"
import "encoding/hex"
import "fmt"
import "io/ioutil"
import "log"
import "path/filepath"

// Challenge4Main solves set 1 challenge 4
func Challenge4Main() {
    fileLocation, err := filepath.Abs("set1/files/4.txt")
    if err != nil { log.Fatal(err) }

    fileBytes, err := ioutil.ReadFile(fileLocation)
    if err != nil { log.Fatal(err) }

    splitLines := bytes.Split(fileBytes, []byte("\n"))

    topScore, buestGuess := 0, []byte(nil)
    for i := 0; i < len(splitLines); i++ {
        hexString := splitLines[i]

        rawBytes := make([]byte, hex.DecodedLen(len(hexString)))
        _, err := hex.Decode(rawBytes, hexString)
        if err != nil { log.Fatal(err) }

        guess, _, score := BreakEncodedText(rawBytes)
        if score > topScore {
            buestGuess = guess
            topScore = score
        }
    }

    fmt.Printf("%s", buestGuess)
}
