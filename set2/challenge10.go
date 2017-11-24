package set2

import "bytes"
import "crypto/aes"
import "encoding/base64"
import "errors"
import "fmt"
import "io/ioutil"
import "log"
import "path/filepath"
import "strings"

import "github.com/AntonKueltz/go-cryptopals/set1"

// DecryptAesCbc Decrypt a ciphertext in AES-CBC
func DecryptAesCbc(key, ciphertext, iv []byte) ([]byte, error) {
    plaintext := make([]byte, len(ciphertext))

    block, err := aes.NewCipher(key)
    if err != nil { log.Fatal(err) }

    blocksize := block.BlockSize()
    if len(ciphertext) % blocksize != 0 {
        return plaintext, errors.New("Ciphertext not multiple of block size")
    }

    for i := 0; i < len(ciphertext) / blocksize; i++ {
        start, end := i * blocksize, (i + 1) * blocksize

        intermediate := make([]byte, blocksize)
        block.Decrypt(intermediate, ciphertext[start:end])
        xored, _ := set1.Xor(intermediate, iv)

        copy(plaintext[start:end], xored)
        copy(iv, ciphertext[start:end])
    }

    return plaintext, nil
}

// Challenge10Main solves set 2 challenge 10
func Challenge10Main() {
    fileLocation, err := filepath.Abs("set2/files/10.txt")
    if err != nil { log.Fatal(err) }

    fileBytes, err := ioutil.ReadFile(fileLocation)
    if err != nil { log.Fatal(err) }
    fileBytes = bytes.Replace(fileBytes, []byte("\n"), []byte(""), -1)

    rawBytes, err := base64.StdEncoding.DecodeString(string(fileBytes))
    if err != nil { log.Fatal(err) }

    key := []byte("YELLOW SUBMARINE")
    iv := []byte(strings.Repeat("\x00", 16))

    result, _ := DecryptAesCbc(key, rawBytes, iv)
    fmt.Printf("%s\n", result)
}
