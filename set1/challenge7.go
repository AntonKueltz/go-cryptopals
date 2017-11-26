package set1

import "bytes"
import "crypto/aes"
import "encoding/base64"
import "errors"
import "fmt"
import "io/ioutil"
import "log"
import "path/filepath"

func padPkcs7(bytes []byte, blocksize int) []byte {
    byteLen := len(bytes)
    paddingBytes := blocksize - (byteLen % blocksize)
    if paddingBytes == 0 {
        paddingBytes = blocksize
    }

    paddedBytes := make([]byte, byteLen + paddingBytes)
    copy(paddedBytes[:byteLen], bytes[:])

    for i := byteLen; i < byteLen + paddingBytes; i++ {
        paddedBytes[i] = byte(paddingBytes)
    }

    return paddedBytes
}

// EncryptAesEcb Encrypt a ciphertext in AES-ECB (not supported by go)
func EncryptAesEcb(key, plaintext []byte) []byte {
    block, err := aes.NewCipher(key)
    if err != nil { log.Fatal(err) }

    blocksize := block.BlockSize()
    plaintext = padPkcs7(plaintext, blocksize)
    ciphertext := make([]byte, len(plaintext))

    for i := 0; i < len(plaintext) / blocksize; i++ {
        start, end := i * blocksize, (i + 1) * blocksize
        block.Encrypt(ciphertext[start:end], plaintext[start:end])
    }

    return ciphertext
}

// DecryptAesEcb Decrypt a ciphertext in AES-ECB (not supported by go)
func DecryptAesEcb(key, ciphertext []byte) ([]byte, error) {
    plaintext := make([]byte, len(ciphertext))

    block, err := aes.NewCipher(key)
    if err != nil { log.Fatal(err) }

    blocksize := block.BlockSize()
    if len(ciphertext) % blocksize != 0 {
        return plaintext, errors.New("Ciphertext not multiple of block size")
    }

    for i := 0; i < len(ciphertext) / blocksize; i++ {
        start, end := i * blocksize, (i + 1) * blocksize
        block.Decrypt(plaintext[start:end], ciphertext[start:end])
    }

    return plaintext, nil
}

// Challenge7Main solves set 1 challenge 7
func Challenge7Main() {
    fileLocation, err := filepath.Abs("set1/files/7.txt")
    if err != nil { log.Fatal(err) }

    fileBytes, err := ioutil.ReadFile(fileLocation)
    if err != nil { log.Fatal(err) }
    fileBytes = bytes.Replace(fileBytes, []byte("\n"), []byte(""), -1)

    rawBytes, err := base64.StdEncoding.DecodeString(string(fileBytes))
    if err != nil { log.Fatal(err) }

    key := []byte("YELLOW SUBMARINE")
    plaintext, _ := DecryptAesEcb(key, rawBytes);
    fmt.Printf("%s\n", plaintext)
}
