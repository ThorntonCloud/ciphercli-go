# CipherCLI - AES Encryption CLI Tool
`CipherCLI`is a simple CLI tool that allows you to encrypt and decrypt text using AES encryption in CFB mode. The secret key is securely loaded from an environment variable, and the encrypted output is base64-encoded for safe storage and transmission.

## Features
- AES encryption in CFB mode with a random IV.
- Base64 encoding for encrypted data.
- CLI Usage for easy integration into shell scripts or terminal commands.
- Supports both encryption and decryption of text.

## Prerequisites
- **Go** version 1.23 or higher installed.
- A `.env` file that contains the secret key (see setup instructions below)

## Installation
### 1. Clone the Repository
```bash
git clone https://github.com/thorntoncloud/ciphercli-go.git
cd ciphercli-go
```
### 2. Setup the `.env` File
Create a `.env` file in the Encrypt directory and add the following line to store your secret key. The key must be 16, 24, or 32 bytes long.
```bash
SECRET_SAUCE=your-secret-key-here
```
Replace `your-secret-key-here` with your actual secret key.

### 3. Build the Executable
To build the tool into an executable, run the following Go command from the Encrypt directory:
```bash
go build -o ciphercli
```
This will create an `ciphercli` executable (or ciphercli.exe on Windows).

You can also name the executable something else, if you like. Just change `ciphercli` after the `-o` flag in the build command to the name you like better.

## Usage
### Encrypting Text
To encrypt text, use the following command:
```bash
./ciphercli -action encrypt -text "your secret message"
```

Example:
```bash
./ciphercli -action encrypt -text "This is a top secret message."
```
The tool will return the base64-encoded encrypted text:
```bash
Encrypted Text: C2FuZG... (base64 output)
```

### Decrypting Text
To decrypt an encrypted text, use the following command:
```bash
./ciphercli -action decrypt -text "base64-encoded-encrypted-text"
```

Example:
```bash
./ciphercli -action decrypt -text "C2FuZG..."
```

The tool will return the original decrypted text:
```bash
Decrypted Text: This is a top secret message
```

### Command-Line Options
- `-action`: Specify the action as either `encrypt` or `decrypt`.
- `-text`: The text to be encrypted or decrypted.

## Adding the Tool to Your PATH
To be able to call the tool from anywhere in the CLI without specifying the full path to the executable, you can add the location of the `ciphercli` binary to your system's `PATH`.

### On Linux / macOS
1. Move the binary to `/usr/local/bin` (or another directory in your `PATH`):
   ```bash
   sudo mv ciphercli /usr/local/bin
   ```
2. Verify it works by typing:
   ```bash
   ciphercli -action encrypt -text "test"
   ```

### On Windows
1. Move the `ciphercli.exe` binary to the desired directory, for example, `C:\tools\ciphercli`.
2. Add the directory to your system's `PATH`:
   - Open **Control Panel**  → **System** → **Advanced System Settings**.
   - Click **Environment Variables**.
   - Under **System Variables**, find the `Path` variable, select it, and click **Edit**.
   - Add the path to the directory where you saved `ciphercli.exe` (e.g., `C:\tools\ciphercli).
   - Click **OK** to save and apply.
3. Verify it works by typing the following command into **Command Prompt** or **PowerShell**
    ```bash
    ciphercli -action encrypt -text "text"
    ```

## Contributing
Feel free to fork this repository, submit pull requests, or file issues with any bug reports or feature requests.

## License
This project is licensed under the MIT License.
