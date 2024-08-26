package handler

import (
	"bufio"
	"errors"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Function to access the banner file
func OpenBanner(banner string, asciiChars map[int][]string) error {
	filename := filepath.Join("./banner", banner+".txt")
	file, err := os.Open(filename)
	// fmt.Println(file)
	if err != nil {
		return err
	}
	defer file.Close()

	scanned := bufio.NewScanner(file)

	code := 31
	for scanned.Scan() {
		line := scanned.Text()
		if line == "" {
			code++
		} else {
			asciiChars[code] = append(asciiChars[code], line)
		}
	}
	if err := scanned.Err(); err != nil {
		log.Fatal(err)
	}
	return nil
}

// Handling the ascii character
func PrintCharAscii(asciiChar map[int][]string, userInput string) (string, error) {
	var result strings.Builder

	lines := strings.Split(userInput, "\r\n")
	for _, line := range lines {
		for j := 0; j < len(asciiChar[32]); j++ {
			for _, letter := range line {
				if letter < 32 || letter > 126 {
					return "", errors.New("Error")
				} else {
					result.WriteString(asciiChar[int(letter)][j])
				}
			}
			result.WriteString("\r\n")
		}
	}
	return result.String(), nil
}

// Function who print the characters ascii art
func PrintAsciiArt(text, banner string) (string, error) {
	asciiChars := make(map[int][]string)
	err := OpenBanner(banner, asciiChars)
	if err != nil {
		return "", err
	}
	return PrintCharAscii(asciiChars, text)
}
