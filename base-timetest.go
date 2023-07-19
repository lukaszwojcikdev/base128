package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Wprowadz ciag znakow:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	// Kodowanie ANSI X3.58
	start := time.Now()
	ansi58Encoded := ansi58Encode(input)
	elapsed := time.Since(start)
	fmt.Printf("Czas kodowania ANSI X3.58: %s\n", elapsed)
	fmt.Println(ansi58Encoded)

	// Dekodowanie ANSI X3.58
	start = time.Now()
	ansi58Decoded := ansi58Decode(ansi58Encoded)
	elapsed = time.Since(start)
	fmt.Printf("Czas dekodowania ANSI X3.58: %s\n", elapsed)
	fmt.Println(ansi58Decoded)

	// Kodowanie Base64
	start = time.Now()
	base64Encoded := base64Encode(input)
	elapsed = time.Since(start)
	fmt.Printf("Czas kodowania Base64: %s\n", elapsed)
	fmt.Println(base64Encoded)

	// Dekodowanie Base64
	start = time.Now()
	base64Decoded := base64Decode(base64Encoded)
	elapsed = time.Since(start)
	fmt.Printf("Czas dekodowania Base64: %s\n", elapsed)
	fmt.Println(base64Decoded)

	// Kodowanie Base128
	start = time.Now()
	base128Encoded := base128Encode(input)
	elapsed = time.Since(start)
	fmt.Printf("Czas kodowania Base128: %s\n", elapsed)
	fmt.Println(base128Encoded)

	// Dekodowanie Base128
	start = time.Now()
	base128Decoded := base128Decode(base128Encoded)
	elapsed = time.Since(start)
	fmt.Printf("Czas dekodowania Base128: %s\n", elapsed)
	fmt.Println(base128Decoded)
}

func ansi58Encode(input string) string {
	encoded := ""
	for _, char := range input {
		if char >= ' ' && char <= '~' { // Zakres znakow widocznych w ASCII
			encoded += string(char)
		}
	}
	return encoded
}

func ansi58Decode(input string) string {
	decoded := ""
	for _, char := range input {
		if char >= ' ' && char <= '~' { // Zakres znakow widocznych w ASCII
			decoded += string(char)
		}
	}
	return decoded
}

func base64Encode(input string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(input))
	return encoded
}

func base64Decode(input string) string {
	decoded, _ := base64.StdEncoding.DecodeString(input)
	return string(decoded)
}

func base128Encode(input string) string {
	encoded := ""
	for _, char := range input {
		encoded += string(char + 1)
	}
	return encoded
}

func base128Decode(input string) string {
	decoded := ""
	for _, char := range input {
		decoded += string(char - 1)
	}
	return decoded
}
