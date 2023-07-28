/******************************************************************************************

Base128 -> Encode and Decode text files
(c) by Łukasz Wójcik 2023
www.base128.com

BUILD:
go build base128.go

RUN:
go run base128.go -h [for HELP]

USE:

	Linux:
		Encode:
			./base128 -e file_encode.txt file_decode.b128
		Decode:
			./base128 -d file_decode.b128 file_encode.txt
			
	Windows:
		Encode:
			base128.exe -e file_encode.txt file_decode.b128
		Decode:
			base128.exe -d file_decode.b128 file_encode.txt
	
LICENSE:
MIT LICENSE
https://mit-license.org

******************************************************************************************/
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

//We define a "base" constant of 128, which defines the encoding range.
const base = 128

//Function to encode data using Base128 data.
//Takes a text string as an argument and returns a sequence of integers representing it.
//Each character from the text is converted to the corresponding number in the range 0-127.
func encodeBase128(input string) []int {
	result := []int{}
	for _, char := range input {
		result = append(result, int(char))
	}
	return result
}

//Function to decode Base128 data.
//Takes a sequence of integers as an argument and converts them to raw text.
//Each number is converted to its corresponding letter or special character.
func decodeBase128(input []int) string {
	result := ""
	for _, code := range input {
		result += string(rune(code))
	}
	return result
}

//The main function of the program that contains the logic of the program.
//Read program call flags, such as "-e" for encoding or "-d" for decoding.
//Then it reads input from 'file' or standard input and processes it with the appropriate function.
//The result is displayed on standard output.
func main() {
	encodeFlag := flag.Bool("e", false, "Converts the input's base128 encoding into an output text file.")
	decodeFlag := flag.Bool("d", false, "Recovers the original input file by decoding the information that was previously encoded using base128.")
	helpFlag := flag.Bool("h", false, "Print instructions for calling and a list of available alternatives.")
	versionFlag := flag.Bool("version", false, "Print the program's version.")
	copyrightFlag := flag.Bool("copyright", false, "Print copyright information.")
	flag.Parse()

	//This snippet displays information about the base128 program, the invocation instructions, and the available options. 
	//If the '*helpFlag' is set to "true", the code will display the appropriate messages and then terminate the program.
	if *helpFlag {
		fmt.Println("base128 - Encodes or decodes FILE, or standard input, to standard output or a file as base128.")
		fmt.Println("base128 [OPTION]... [FILE input] > [FILE output]")
		fmt.Println()
		fmt.Println("EXAMPLE encode: base128 -e encode.txt > decode.txt")
		fmt.Println("EXAMPLE decode: base128 -d decode.txt > encode2.txt")
		fmt.Println()
		fmt.Println("Options:")
		fmt.Println()
		fmt.Println("  -d, --decode          Recovers the original input file by decoding the information that was previously encoded using base128.")
		fmt.Println("  -e, --encode          Converts the input's base128 encoding into an output text file.")
		fmt.Println("  -h, --help            Print instructions for calling and a list of available alternatives.")
		fmt.Println("      --version         Print the program's version.")
		fmt.Println("      --copyright       Print copyright information.")
		fmt.Println()
		fmt.Println("(c) by Lukasz Wojcik 2023")
		fmt.Println("http://www.base128.com")
		os.Exit(0)
	}

	//Displays information about the current version of base128, then exits the program.
	if *versionFlag {
		fmt.Println("base128 version 1.0/2023")
		os.Exit(0)
	}
	
	//Displays base128's copyright and program author information, then exits the program.
	if *copyrightFlag {
		fmt.Println("base128 (C) 2023 by Lukasz Wojcik")
		os.Exit(0)
	}

	//Checks if both *encodeFlag and *decodeFlag are set to true at the same time. 
	//If so, it prints an error message on stdout using fmt.Fprintln using os.Stderr and exits with an error code of 1. 
	//The message informs the user that he cannot specify a flag for both encoding and decoding at the same time, and suggests using the flag help (*helpFlag) for more information.
	if *encodeFlag && *decodeFlag {
		fmt.Fprintln(os.Stderr, "base128: cannot specify. Please -e for ENCODE or -d for DECODE or -h for HELP")
		os.Exit(1)
	}

	//This snippet checks that neither of the *encodeFlag and *decodeFlag flags are set to true. 
	//If so, the code prints an error message to stdout using fmt.Fprintln and os.Stderr and exits the program with an error code of 1. 
	//The message informs the user that he has not specified any flag for encoding or decoding and suggests the use of a flag help (*helpFlag) for more information.
	if !*encodeFlag && !*decodeFlag {
		fmt.Fprintln(os.Stderr, "base128: cannot specify. Please -e for ENCODE or -d for DECODE or -h for HELP")
		os.Exit(1)
	}

	args := flag.Args()
	var input string
	if len(args) == 0 || args[0] == "-" {
		bytes, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}
		input = strings.TrimSpace(string(bytes))
	} else {
		bytes, err := ioutil.ReadFile(args[0])
		if err != nil {
			panic(err)
		}
		input = strings.TrimSpace(string(bytes))
	}

	if *encodeFlag {
		input = strings.Replace(input, "\n", "", -1)
		input = strings.Replace(input, "\r", "", -1)
		input = strings.TrimSpace(input)
		numbers := encodeBase128(input)
		output := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(numbers)), " "), "[]")
		fmt.Println(output)
	} else if *decodeFlag {
		input = strings.Replace(input, "\n", "", -1)
		input = strings.Replace(input, "\r", "", -1)
		input = strings.TrimSpace(input)
		numbersStr := strings.Split(input, " ")
		numbers := make([]int, len(numbersStr))
		for i, numberStr := range numbersStr {
			numberStr = strings.TrimSpace(numberStr)
			number, err := strconv.Atoi(numberStr)
			if err != nil {
				panic(err)
			}
			numbers[i] = number
		}
		output := decodeBase128(numbers)
		fmt.Println(output)
	}
 }
