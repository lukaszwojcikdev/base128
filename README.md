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
![version](https://img.shields.io/badge/version-1.0-blue)
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
 - [Download](#download)
 - [License](#license)

## Application

Here are some examples of uses in the context of base128 security:

- Data hiding:
    - Base128 can be used to hide data that is not understandable to the naked eye or when intercepted by a third party.
    - By encoding data with Base128, unauthorized access to sensitive information can be prevented, which is of great importance in data security.
- Obstruction of reading and manipulation:
    - Base128 makes it difficult for unauthorized persons to read and manipulate data because it is less intuitive than other popular encoding formats such as Base64.
    - This means that if the data is intercepted by unauthorized persons, it will be more difficult for them to understand and modify this data.
- Data transport:
    - Can be used to securely transport data between various system components.
    - This includes the secure transmission of passwords, authentication tokens and other sensitive information over a network where undesirable reading or alteration of data is a threat.
- Pseudonymization of data:
    - Serves as a data pseudonymization tool, allowing access to data presented in a different, irreversible format.
    - This can be used in the context of data analysis and protecting sensitive information from unauthorized access.
      
## Demo (in production)

![App Screenshot](https://via.placeholder.com/268x150?text=App+Screenshot+Here)


## Instalation

To install Base128, follow these steps:

- Clone repository from GitHub:

``` git clone https://github.com/lukaszwojcikdev/base128.git ``` 

- Go to the project directory: ``` cd base128 ``` 

- Compile the source code: ``` go build base128.go ```

- Ready! The program was compiled for Windows as **base128.exe** for Linux **./base128**
   
- Base128 is now ready to use.
  
   
## Usage/Examples

```javascript
NAME

base128 - Encode and Decode text files

SYNOPSIS
base128 [ -e or -d ]
[ options ] [ input file ] > [ output file ]
[ example encode: ] base128 -e [ encode.txt ] > [ decode.txt ]
[ example decode: ] base128 -d [ decode.txt ] > [ encode.txt ]

base128 is a command line tool that encodes and decodes text files, e.g. *.txt , *.svg , *.html

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

#### Coding ####
Suppose we want to encode the file "data.txt" and save the encoded data to the file "encode.bin".

In this situation, we execute the command:

``` ./base128 -e data.txt encode.bin ``` 
or
``` ./base128 -e data.txt >> encode.bin ``` 

#### Decoding ####
Suppose we have a file "encoded.bin" containing data encoded in Base128 format and we want to decode it into a file "decoded.txt".

In this case, we invoke the command:

``` ./base128 -d encoded.bin decoded.txt ``` 
or

``` ./base128 -d encoded.bin >> decoded.txt ``` 

That's all!

Now you should be able to install Base128 and use it to encode and decode data using the Base128 format.

## Contributing

Contributions are always welcome!

See `contributing.md` for ways to get started.

Please adhere to this project's `code of conduct`.


## Documentation

[See these PDF](http://www.base128.com/base128.com/base128.pdf)


## Author

- [@lukaszwojcikdev](https://www.github.com/lukaszwojcikdev)


## Site

- [www.base128.com](http://www.base128.com)
- [www.lukaszwojcik.eu](http://www.lukaszwojcik.eu)
## Download

Windows|Linux
-|-
[ZIP](http://www.base128.com/base128.com/base128.zip)|[TAR.GZ](http://www.base128.com/base128.com/base128.tar.gz)
[MD5](http://www.base128.com/base128/base128.md5sum/)|[MD5](http://www.base128.com/base128.com/base128.tar.gz.md5sum)
## License

[MIT License](https://choosealicense.com/licenses/mit/)

