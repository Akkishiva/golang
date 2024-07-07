package main

import (
    "fmt"
    "unicode/utf8"
)

func main() {
    // A string containing Unicode characters
    str := "Hello, 世界"  // "世界" means "world" in Chinese

    // Print the string
    fmt.Println("Original String:", str)

    // Iterate over each rune in the string
    fmt.Println("Runes in the string:")
    for _, r := range str {
        fmt.Printf("%c ", r)
    }
    fmt.Println()

    // Length of the string in bytes
    fmt.Println("Length in bytes:", len(str))

    // Length of the string in runes
    fmt.Println("Length in runes:", utf8.RuneCountInString(str))
}
