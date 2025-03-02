```
  __________                          ____ ________    ______  
  \______   \_____     ______  ____  /_   |\_____  \  /  __  \ 
   |    |  _/\__  \   /  ___/_/ __ \  |   | /  ____/  >      < 
   |    |   \ / __ \_ \___ \ \  ___/  |   |/       \ /   --   \
   |______  /(____  //____  > \___  > |___|\_______ \\______  /
          \/      \/      \/      \/               \/       \/
```
# Base128

[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/) 
![Website](https://img.shields.io/website?url=http%3A%2F%2Fbase128.com)
![version](https://img.shields.io/badge/version-2.0-blue)
![Golang](https://img.shields.io/badge/-Golang-00ADD8?logo=Go&logoColor=white&style=flat)
[![LinkedIn](https://img.shields.io/badge/LinkedIn-Connect-blue?style=social&logo=linkedin)](https://www.linkedin.com/in/lukasz-michal-wojcik)

> OpenSource tool for encoding and decoding text data.

It works by numerically representing the input data with 7-bit groups of bits, which enables efficient data storage and transfer.

## Contents

 - [Application](#application)
 - [Instalation](#instalation)
 - [Usage/Examples](#usageexamples)
 - [Contributing](#contributing)
 - [Documentation](#documentation)
 - [Autor](#author)
 - [Site](#site)
 - [Donation](#donation)
 - [Download](#download)
 - [License](#license)

## Application

Here are some examples of uses in the context of **Base128 text data transformation**:

### **Text Data Transformation**
- **Base128** can be used to transform text data into a sequence of integers, making it less human-readable.
- This can be useful for scenarios where text data needs to be represented in a numeric format, such as storing text as a list of integers in a database.

### **Educational Purposes**
- The program serves as a simple tool to demonstrate basic encoding and decoding concepts.
- It can be used to teach how text characters are mapped to their ASCII values and vice versa.

### **Data Representation**
- **Base128** can be used to represent text data in a numeric format, which may be useful in certain data processing or storage scenarios.
- Example: Encoding a message like `"Hello"` into `72 101 108 108 111`.

## **Limitations**
- **Text-Only Support**: The program is designed to work with **text data** and does not support binary data. Attempting to encode or decode binary data may result in errors or incorrect output.
- **No Compression or Security**: The algorithm does not compress data or provide encryption. It simply converts text to a sequence of integers and vice versa.
- **Limited Character Range**: The algorithm works with characters in the range 0-127 (ASCII), which means it cannot handle extended character sets or binary data.

## **Demo (in production)**

![App Screenshot](https://via.placeholder.com/268x150?text=App+Screenshot+Here)

## **Installation**

To install **Base128**, follow these steps:

1. **Clone the repository from GitHub**:
   ```bash
   git clone https://github.com/lukaszwojcikdev/base128.git
   ```

2. **Go to the project directory**:
   ```bash
   cd base128
   ```

3. **Compile the source code**:
   ```bash
   go build base128.go
   ```

4. **Ready!** The program is now compiled:
   - For Windows: **`base128.exe`**
   - For Linux: **`./base128`**

5. **Base128** is now ready to use.
  
   
## Usage/Examples

```javascript
  __________                          ____ ________    ______
  \______   \_____     ______  ____  /_   |\_____  \  /  __  \
   |    |  _/\__  \   /  ___/_/ __ \  |   | /  ____/  >      <
   |    |   \ / __ \_ \___ \ \  ___/  |   |/       \ /   --   \
   |______  /(____  //____  > \___  > |___|\_______ \\______  /
          \/      \/      \/      \/               \/       \/
NAME
base128 - encodes or decodes a text file on standard input and standard output
SYNOPSIS
base128 [ -e or -d ]
[ options ] [ input file ] > [ output file ]
[ example encode: ] base128 -e [ encode_text.txt ] > [ decode_text.txt ]
[ example decode: ] base128 -d [ decode_text.txt ] > [ encode2_text.txt ]
base128 Is a command line tool that encodes and decodes text files, e.g. *.txt , *.svg , *.html

OPTIONS
-e, --encode
Converts the input's base128 encoding into an output text file.
-d, --decode
Recovers the original input file by decoding the information that was previously encoded using base128.
-h, --help
Print instructions for calling and a list of available alternatives.
--version
Print the program's version.
--copyright
Print copyright information.
```

## **Usage**

### **Encoding**
Suppose we want to encode the file `data.txt` and save the encoded data to the file `encode.txt`.

In this situation, execute the command:
```bash
./base128 -e data.txt > encode.txt
```

### **Decoding**
Suppose we have a file `encoded.txt` containing data encoded in Base128 format, and we want to decode it into a file `decoded.txt`.

In this case, invoke the command:
```bash
./base128 -d encoded.txt > decoded.txt
```

That's all! Now you should be able to install **Base128** and use it to encode and decode text data using the Base128 format.


## **Contributing**

Contributions are always welcome! If you'd like to contribute to the development of **Base128**, please follow these steps:

1. Fork the repository on GitHub.
2. Create a new branch for your feature or bugfix.
3. Make your changes and ensure the code is properly tested.
4. Submit a pull request with a detailed description of your changes.


## Documentation

[See these PDF](http://www.base128.pl/base128_source_code.pdf)


## Author

- [@lukaszwojcikdev](https://www.github.com/lukaszwojcikdev)


## Site

- [www.base128.pl](http://www.base128.pl)
- [www.lukaszwojcik.eu](http://www.lukaszwojcik.eu)
  
## Download

Windows|Linux
-|-
[ZIP](http://www.base128.pl/base128.zip)|
[MD5](http://www.base128.pl/base128.md5sum/)|
## License

[MIT License](https://choosealicense.com/licenses/mit/)

