package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const (
	lowerCharSet   = "abcdefghijklmnopqrstuvwxyz"
	upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberCharSet  = "0123456789"
	specialCharSet = "!@#$%^&*()_+-=[]{}|;:,.<>/?`~"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var length int
	var includeLowercase, includeUppercase, includeNumbers, includeSpecialChars string

	// Validate password length
	for {
		fmt.Print("Enter the desired password length: ")
		var input string
		fmt.Scan(&input)
		var err error
		length, err = strconv.Atoi(input)
		if err != nil || length <= 0 {
			fmt.Println("Please enter a valid positive integer for the password length.")
		} else {
			break
		}
	}

	// Validate include lowercase characters
	includeLowercase = getValidInput("Include lowercase characters? (y/n): ")

	// Validate include uppercase characters
	includeUppercase = getValidInput("Include uppercase characters? (y/n): ")

	// Validate include numbers
	includeNumbers = getValidInput("Include numbers? (y/n): ")

	// Validate include special characters
	includeSpecialChars = getValidInput("Include special characters? (y/n): ")

	lowercase := includeLowercase == "y"
	uppercase := includeUppercase == "y"
	numbers := includeNumbers == "y"
	specialChars := includeSpecialChars == "y"

	password := generatePassword(length, lowercase, uppercase, numbers, specialChars)
	fmt.Printf("Generated password: %s\n", password)

	// Wait for user input before closing
	fmt.Print("Press Enter to exit...")
	fmt.Scanln()
	fmt.Scanln() // Extra Scanln to handle Enter key press after input
}

func getValidInput(prompt string) string {
	for {
		fmt.Print(prompt)
		var input string
		fmt.Scan(&input)
		input = strings.ToLower(input)
		if input == "y" || input == "n" {
			return input
		}
		fmt.Println("Please enter 'y' or 'n'.")
	}
}

func generatePassword(length int, includeLowercase, includeUppercase, includeNumbers, includeSpecialChars bool) string {
	var charSets string

	if includeLowercase {
		charSets += lowerCharSet
	}

	if includeUppercase {
		charSets += upperCharSet
	}

	if includeNumbers {
		charSets += numberCharSet
	}

	if includeSpecialChars {
		charSets += specialCharSet
	}

	if len(charSets) == 0 {
		return ""
	}

	var password []byte
	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(len(charSets))
		password = append(password, charSets[randomIndex])
	}

	return string(password)
}
