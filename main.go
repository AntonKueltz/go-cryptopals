package main

import "flag"
import "fmt"

import "github.com/AntonKueltz/go-cryptopals/set1"
import "github.com/AntonKueltz/go-cryptopals/set2"

func main() {
    numbPtr := flag.Int("set", 0, "Which set to run (0=all)")
    flag.Parse()

    if *numbPtr == 0 || *numbPtr == 1 {
        fmt.Println("[Set 1] Challenge 1")
        set1.Challenge1Main()

        fmt.Println("\n[Set 1] Challenge 2")
        set1.Challenge2Main()

        fmt.Println("\n[Set 1] Challenge 3")
        set1.Challenge3Main()

        fmt.Println("\n[Set 1] Challenge 4")
        set1.Challenge4Main()

        fmt.Println("\n[Set 1] Challenge 5")
        set1.Challenge5Main()

        fmt.Println("\n[Set 1] Challenge 6")
        set1.Challenge6Main()

        fmt.Println("\n[Set 1] Challenge 7")
        set1.Challenge7Main()

        fmt.Println("\n[Set 1] Challenge 8")
        set1.Challenge8Main()
        fmt.Println()
    }

    if *numbPtr == 0 || *numbPtr == 2 {
        fmt.Println("[Set 2] Challenge 9")
        set2.Challenge9Main()
    }
}
