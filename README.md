# CipherCLI - AES Encryption CLI Tool
`ciphercli` is a command-line tool for securely encrypting and decrypting text using AES encryption in CFB mode. The tool loads a secret key from an environment variable, ensuring secure handling of sensitive data. Encrypted output is base64-encoded, making it suitable for storage or transmission.

## Features
- AES encryption in CFB mode with a random IV for enhanced security.
- Base64 encoding for safe storage of encrypted data.
- Command-line interface with simple, clear subcommands for encryption and decryption.
- Ideal for integration into scripts or standalone use in the terminal.

## Prerequisites
- Go version 1.23 or higher.
- `.env` file containing the secret key (see setup instructions below).

## Installation
### 1. Clone the Repository
```bash
git clone https://github.com/thorntoncloud/ciphercli-go.git
cd ciphercli-go
```

### 2. Set Up the .env File
Create a `.env` file in the project’s root directory and add your secret key (16, 24, or 32 bytes long):

```bash
SECRET_SAUCE=your-secret-key-here
```
Replace your-secret-key-here with your actual key.

### 3. Build the Executable
To build the CLI tool, run:
```bash
go build -o ciphercli
```
This will create a `ciphercli` executable (or `ciphercli.exe` on Windows) in the project directory.

## Usage
CipherCLI uses spf13/cobra for a clear, structured command interface. You can choose encrypt or decrypt subcommands with a specified text to encrypt or decrypt.

### Encrypt Text
To encrypt text, use the encrypt command:

```bash
./ciphercli encrypt "your secret message"
```
Example:
```bash
./ciphercli encrypt "This is a top secret message."
```

The CLI will return the base64-encoded encrypted text:

```bash
Encrypted Text: C2FuZG... (base64 output)
```
### Decrypt Text
To decrypt encrypted text, use the decrypt command:

```bash
./ciphercli decrypt "base64-encoded-encrypted-text"
```
Example:
```bash
./ciphercli decrypt "C2FuZG..."
```

The CLI will return the original decrypted text:

```bash
Decrypted Text: This is a top secret message
```

### Command-Line Options
- `encrypt`: Encrypts the specified text.
- `decrypt`: Decrypts the specified base64-encoded encrypted text.

Example commands:

```bash
ciphercli encrypt "Text to encrypt"
ciphercli decrypt "EncryptedTextHere"
```

## Adding the Tool to Your PATH
You can add `ciphercli` to your PATH for easier access.

### On Linux / macOS
Move the binary to /usr/local/bin:

```bash
sudo mv ciphercli /usr/local/bin
```

Verify it’s accessible:
```bash
ciphercli encrypt "test message"
```

### On Windows
Move `ciphercli.exe` to your preferred directory, e.g., C:\tools\ciphercli.

Add the directory to your system’s PATH:

1. Open Control Panel → System → Advanced System Settings.
2. Go to Environment Variables.
3. Under System Variables, find the Path variable, select it, and click Edit.
4. Add the path to the directory containing ciphercli.exe (e.g., C:\tools\ciphercli).
5. Click OK to save.

Verify it’s accessible:
```bash
ciphercli encrypt "text"
```

## Contributing
Contributions are welcome! Feel free to fork this repository, open pull requests, or file issues for bug reports or feature requests.

## License
This project is licensed under the MIT License.