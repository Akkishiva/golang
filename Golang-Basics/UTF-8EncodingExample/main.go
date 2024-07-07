package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// Prompt user to enter a string
	fmt.Print("Enter a string: ")
	var input string
	fmt.Scanln(&input)

	// Print the input string
	fmt.Printf("Input String: %s\n", input)

	// Iterate over each character in the string
	fmt.Println("UTF-8 Encodings:")
	for _, runeValue := range input {
		// Get the UTF-8 encoding for each character
		utf8Bytes := make([]byte, utf8.RuneLen(runeValue))
		utf8.EncodeRune(utf8Bytes, runeValue)

		// Print the UTF-8 encoding in binary and hexadecimal formats
		fmt.Printf("%s: ", string(runeValue))

		// Print binary representation
		fmt.Printf("Binary: ")
		for i := 0; i < len(utf8Bytes); i++ {
			fmt.Printf("%08b ", utf8Bytes[i])
		}

		// Print hexadecimal representation
		fmt.Printf("Hexadecimal: ")
		for _, b := range utf8Bytes {
			fmt.Printf("%X ", b)
		}

		fmt.Println()
	}
}
