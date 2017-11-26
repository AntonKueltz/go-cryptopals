package set2

import "crypto/aes"
import "fmt"
import "log"
import "math/rand"
import "net/url"
import "strings"

import "github.com/AntonKueltz/go-cryptopals/set1"

var c13Key = []byte("COUNTERATTACKERS")  // or something else random...

func profileFor(email string) string {
    email = strings.Replace(email, "&", "", -1)
    email = strings.Replace(email, "=", "", -1)

    return fmt.Sprintf("email=%s&uid=%d&role=user", email, rand.Intn(100))
}

func decodeProfile(profile string) map[string][]string {
    keyValue, err := url.ParseQuery(profile)
    if err != nil { log.Fatal(err) }

    return keyValue
}

func encryptProfile(profile string) []byte {
    return set1.EncryptAesEcb(c13Key, []byte(profile))
}

func decryptProfile(encrypted []byte) []byte {
    decrypted, err := set1.DecryptAesEcb(c13Key, encrypted)
    if err != nil { log.Fatal(err) }

    return decrypted
}

// Challenge13Main solves set 2 challenge 13
func Challenge13Main() {
    adminEmail := "AAAAAAAAAAadmin" + strings.Repeat("\x11", 11)
    adminProfile := profileFor(adminEmail)
    encryptedAdmin := encryptProfile(adminProfile)
    adminBlock := encryptedAdmin[aes.BlockSize:(2 * aes.BlockSize)]

    userEmail := "pwned!@pwn.me"
    userProfile := profileFor(userEmail)
    encryptedUser := encryptProfile(userProfile)

    cutPasted := make([]byte, len(encryptedUser))
    beforeAdmin := len(encryptedUser) - aes.BlockSize
    copy(cutPasted[:beforeAdmin], encryptedUser[:beforeAdmin])
    copy(cutPasted[beforeAdmin:], adminBlock)

    decrypted := decryptProfile(cutPasted)
    keyValue := decodeProfile(string(decrypted))
    fmt.Println(keyValue)
}
