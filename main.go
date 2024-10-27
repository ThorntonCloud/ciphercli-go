// main.go
package main

import (
	"flag"
	"fmt"

	"github.com/thorntoncloud/ciphercli-go/config"
	"github.com/thorntoncloud/ciphercli-go/encryption"
)

func main() {
	// Initialize configuration
	config.InitConfig()
	secretKey := config.GetSecret()

	if secretKey == "" {
		fmt.Println("Secret key is missing")
		return
	}

	// Define flags for command-line arguments
	action := flag.String("action", "encrypt", "Specify 'encrypt' or 'decrypt'")
	text := flag.String("text", "", "Text to encrypt or decrypt")

	flag.Parse()

	if *text == "" {
		fmt.Println("Please provide the text using the -text flag.")
		return
	}

	if *action == "encrypt" {
		// Encrypt the text
		encText, err := encryption.Encrypt(*text, secretKey)
		if err != nil {
			fmt.Println("Error encrypting your classified text:", err)
			return
		}
		fmt.Println("Encrypted Text:", encText)
	} else if *action == "decrypt" {
		// Decrypt the text
		decText, err := encryption.Decrypt(*text, secretKey)
		if err != nil {
			fmt.Println("Error decrypting your encrypted text:", err)
			return
		}
		fmt.Println("Decrypted Text:", decText)
	} else {
		fmt.Println("Invalid action. Please specify 'encrypt' or 'decrypt'.")
	}
}
