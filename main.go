// main.go
package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/thorntoncloud/ciphercli-go/config"
	"github.com/thorntoncloud/ciphercli-go/encryption"
)

var secretKey string

var rootCmd = &cobra.Command{
	Use:   "ciphercli",
	Short: "A simple CLI tool for encryption and decryption",
	Long:  `ciphercli is a command-line interface tool that allows you to encrypt and decrypt text using AES encryption.`,
}

var encryptCmd = &cobra.Command{
	Use:   "encrypt [text]",
	Short: "Encrypt the provided text",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		text := args[0]
		encText, err := encryption.Encrypt(text, secretKey)
		if err != nil {
			fmt.Println("Error encrypting your classified text:", err)
			return
		}
		fmt.Println("Encrypted Text:", encText)
	},
}

var decryptCmd = &cobra.Command{
	Use:   "decrypt [text]",
	Short: "Decrypt the provided text",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		text := args[0]
		decText, err := encryption.Decrypt(text, secretKey)
		if err != nil {
			fmt.Println("Error decrypting your encrypted text:", err)
			return
		}
		fmt.Println("Decrypted Text:", decText)
	},
}

var version = "1.0.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display the current version of CipherCLI",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("CipherCLI Version:", version, "🎉")
	},
}

func main() {
	// Initialize configuration
	config.InitConfig()
	secretKey = config.GetSecret()

	if secretKey == "" {
		fmt.Println("Secret key is missing")
		os.Exit(1)
	}

	// Add commands to root command
	rootCmd.AddCommand(encryptCmd, decryptCmd, versionCmd)

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
