package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func isPHPExtensionInstalled(extension string) bool {
	cmd := exec.Command("php", "-m")

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error running command:", err)
		return false
	}

	outputString := strings.ToLower(string(output))
	return strings.Contains(outputString, extension)
}

func printStatus(extension string, installed bool) {
	status := "❌"
	if installed {
		status = "✔"
	}
	fmt.Printf("PHP %s Extension: %s\n", extension, status)
}

func main() {
	extensions := []string{
		"ctype", "curl", "dom", "fileinfo", "filter", "hash", "mbstring",
		"openssl", "pcre", "pdo", "session", "tokenizer", "xml",
	}

	for _, extension := range extensions {
		installed := isPHPExtensionInstalled(extension)
		printStatus(extension, installed)
	}
}
