package set2

import "fmt"

// PadPkcs7 pads a byte array accoring to the PKCS#7 standard
func PadPkcs7(bytes []byte, blocksize int) []byte {
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

// Challenge9Main solves set 2 challenge 9
func Challenge9Main() {
    stringToPad := []byte("YELLOW SUBMARINE")
    padded := PadPkcs7(stringToPad, 20)
    fmt.Printf("%s\n", padded)
    fmt.Println(padded)
}
