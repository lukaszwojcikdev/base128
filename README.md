
![Logo](http://www.base128.com/base128.com/base128.svg)

# Base128

[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/) 
![Website](https://img.shields.io/website?url=http%3A%2F%2Fbase128.com)
![version](https://img.shields.io/badge/version-1.0-blue)
![Golang](https://img.shields.io/badge/-Golang-00ADD8?logo=Go&logoColor=white&style=flat)
[![LinkedIn](https://img.shields.io/badge/LinkedIn-Connect-blue?style=social&logo=linkedin)](https://www.linkedin.com/in/lukasz-michal-wojcik)

> Narzędzie OpenSource do kodowania i dekodowania danych tekstowych.

Działa na zasadzie reprezentacji liczbowej danych wejściowych za pomocą 7-bitowych grup bitów, co umożliwia efektywne przechowywanie i przesyłanie danych.


## Contents

 - [Zastosowanie](#zastosowanie)
 - [Instalacja](#instalacja)
 - [Użycie](#usageexamples)
 - [Wsparcie](#contributing)
 - [Dokumentacja](#documentation)
 - [Autor](#authors)
 - [Strona](#site)
 - [Pobieranie](#download)
 - [Licencja](#license)

## Zastosowanie

Oto kilka przykładów zastosowań w kontekście bezpieczeństwa base128:

- Ukrywanie danych:
   - Base128 może być używany do ukrywania danych, które nie są zrozumiałe dla zwykłego oka lub w przypadku przechwycenia przez osobę trzecią.
   - Kodując dane za pomocą Base128, można zapobiec nieuprawnionemu dostępowi do wrażliwych informacji, co ma duże znaczenie w bezpieczeństwie danych.
- Utrudnianie odczytu i manipulacj:
   - Base128 utrudnia odczytanie i manipulację danych przez niepowołane osoby, ponieważ jest to mniej intuicyjne niż inne popularne formaty kodowania, takie jak Base64.
   - Oznacza to, że jeśli dane zostaną przechwycone przez niepowołane osoby, trudniej będzie im zrozumieć i zmodyfikować te dane.
- Transport danych:
   - Może być używany do bezpiecznego transportu danych między różnymi składnikami systemu.
   - Obejmuje to bezpieczne przesyłanie haseł, tokenów uwierzytelniających oraz innych wrażliwych informacji przez sieć, gdzie niepożądane odczytywanie lub zmienianie danych stanowi zagrożenie.
- Pseudonimizacja danych:
   -  Służy jako narzędzie do pseudonimizacji danych, pozwalając na dostęp do danych przedstawionych w innym, nieodwracalnym formacie.
   -  Można to wykorzystać w kontekście analizy danych i zabezpieczania wrażliwych informacji przed nieuprawnionym dostępem.
   -  
## Demo (in production)

![App Screenshot](https://via.placeholder.com/268x150?text=App+Screenshot+Here)

## Instalacja 

Aby zainstalować Base128, wykonaj następujące kroki: 

- Sklonuj repozytorium z GitHuba:

``` git clone https://github.com/lukaszwojcikdev/base128.git ``` 

- Przejdź do katalogu z projektem: ``` cd base128 ``` 

- Skompiluj kod źródłowy: ``` go build base128.go ```

- Gotowe! Program został skompilowany dla Windows jako **base128.exe** dla Linux **./base128**
   
- Base128 jest teraz gotowe do użycia.
   
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

#### Kodowanie ####
Załóżmy, że chcemy zakodować plik "dane.txt" i zapisać zakodowane dane do pliku "zakodowane.bin". 

W tej sytuacji wykonujemy polecenie: 

``` ./base128 -e dane.txt zakodowane.bin ``` 
lub
``` ./base128 -e dane.txt >> zakodowane.bin ``` 

#### Dekodowanie ####
Załóżmy, że mamy plik "zakodowane.bin" zawierający dane zakodowane w formacie Base128 i chcemy je zdekodować do pliku "odkodowane.txt". 

W tym przypadku wywołujemy polecenie: 

``` ./base128 -d zakodowane.bin odkodowane.txt ``` 
lub

``` ./base128 -d zakodowane.bin >> odkodowane.txt ``` 

To wszystko! 

Teraz powinieneś być w stanie zainstalować Base128 oraz skorzystać z niego do kodowania i dekodowania danych przy użyciu formatu Base128.
## Contributing

Contributions are always welcome!

See `contributing.md` for ways to get started.

Please adhere to this project's `code of conduct`.


## Documentation

[See these PDF](http://www.base128.com/base128.com/base128.pdf)


## Authors

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

