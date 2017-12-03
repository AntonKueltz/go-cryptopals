package set2

import "fmt"

// ValidatePkcs7Padding returns whether bytes are padded according to PKCS#7
func ValidatePkcs7Padding(buffer []byte) bool {
    buflen := len(buffer)
    lastByte := buffer[buflen - 1]

    for i := 0; i < int(lastByte); i++ {
        if buffer[buflen - 1 - i] != lastByte {
            return false
        }
    }

    return true
}

// Challenge15Main solves set 2 challenge 15
func Challenge15Main() {
    tests := map[string]string{
        "ICE ICE BABY\x04\x04\x04\x04": "ICE ICE BABY\\x04\\x04\\x04\\x04",
        "ICE ICE BABY\x05\x05\x05\x05": "ICE ICE BABY\\x05\\x05\\x05\\x05",
        "ICE ICE BABY\x01\x02\x03\x04": "ICE ICE BABY\\x01\\x02\\x03\\x04",
    }

    for key, value := range tests {
        if ValidatePkcs7Padding([]byte(key)) {
            fmt.Println(value + " is PKCS#7 padded")
        } else {
            fmt.Println(value + " is not PKCS#7 padded")
        }
    }
}
