package set2

import "crypto/aes"
import "fmt"
import "log"
import "math/rand"
import "strings"
import "time"

import "github.com/AntonKueltz/go-cryptopals/set1"

// Challenge11Main solves set 2 challenge 11
func Challenge11Main() {
    prepend := make([]byte, 7)
    append := make([]byte, 7)
    key := make([]byte, aes.BlockSize)
    iv := make([]byte, aes.BlockSize)

    _, err := rand.Read(prepend)
    if err != nil { log.Fatal(err) }
    _, err = rand.Read(append)
    if err != nil { log.Fatal(err) }
    _, err = rand.Read(key)
    if err != nil { log.Fatal(err) }
    _, err = rand.Read(iv)
    if err != nil { log.Fatal(err) }

    data := []byte(strings.Repeat("A", 100))
    plaintext := make([]byte, len(prepend) + len(data) + len(append))

    copy(plaintext[:len(prepend)], prepend)
    copy(plaintext[len(prepend):len(prepend) + len(data)], data)
    copy(plaintext[len(data) + len(prepend):], append)

    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    var ciphertext []byte
    if r.Intn(2) == 1 {
        fmt.Println("Encrypting in CBC mode...")
        ciphertext = EncryptAesCbc(key, plaintext, iv)
    } else {
        fmt.Println("Encrypting in ECB mode...")
        ciphertext = set1.EncryptAesEcb(key, plaintext)
    }

    ecbMode := set1.DetectEcbMode(ciphertext)

    if ecbMode {
        fmt.Println("Detected ECB mode")
    } else {
        fmt.Println("Detected CBC mode")
    }
}
