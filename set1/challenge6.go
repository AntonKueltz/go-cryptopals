package set1

import "bytes"
import "encoding/base64"
import "errors"
import "fmt"
import "io/ioutil"
import "log"
import "path/filepath"
import "sort"

func hammingDistance(bytes1, bytes2 []byte) (int, error) {
    distance := 0

    if len(bytes1) != len(bytes2) {
        return distance, errors.New("Inputs must be the same length")
    }

    for i := 0; i < len(bytes1); i++ {
        difference := bytes1[i] ^ bytes2[i]

        for difference != 0 {
            if (difference & 1) == 1 {
                distance++
            }

            difference >>= 1
        }
    }

    return distance, nil
}

func determineKeySize(encoded []byte) int {
    distanceToKeysize := make(map[float64]int)
    var distances []float64

    for keysize := 2; keysize < 40; keysize++ {
        editDistance, blocks := 0, len(encoded) / keysize

        for i := 1; i < blocks; i++ {
            block1 := encoded[((i-1) * keysize):(i * keysize)]
            block2 := encoded[(i * keysize):((i+1) * keysize)]

            distance, _ := hammingDistance(block1, block2)
            editDistance += distance
        }

        normalizedDistance := float64(editDistance) / float64(keysize * blocks)
        distanceToKeysize[normalizedDistance] = keysize
        distances = append(distances, normalizedDistance)
    }

    sort.Float64s(distances)
    return distanceToKeysize[distances[0]]
}

func transposeBlocks(encoded []byte, keysize int) [][]byte {
    blocks := make([][]byte, keysize)

    for i := 0; i < len(encoded); i++ {
        blocks[i % keysize] = append(blocks[i % keysize], encoded[i])
    }

    return blocks
}

func decode(encoded []byte) []byte {
    keysize := determineKeySize(encoded)

    blocks := transposeBlocks(encoded, keysize)
    key := make([]byte, keysize)

    for i, block := range blocks {
        _, keyByte, _ := BreakEncodedText(block)
        key[i] = keyByte
    }

    fmt.Printf("Key: %s\n\n", key)
    return RepeatingKeyXor(key, encoded)
}

// Challenge6Main solves set 1 challenge 6
func Challenge6Main() {
    fileLocation, err := filepath.Abs("set1/files/6.txt")
    if err != nil { log.Fatal(err) }

    fileBytes, err := ioutil.ReadFile(fileLocation)
    if err != nil { log.Fatal(err) }
    fileBytes = bytes.Replace(fileBytes, []byte("\n"), []byte(""), -1)

    rawBytes, err := base64.StdEncoding.DecodeString(string(fileBytes))
    if err != nil { log.Fatal(err) }

    decoded := decode(rawBytes)
    fmt.Printf("%s\n", decoded)
}
