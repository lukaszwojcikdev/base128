# Base128 jest narzędziem służącym do kodowania i dekodowania danych w formacie Base128. #
Działa na zasadzie reprezentacji liczbowej danych wejściowych za pomocą 7-bitowych grup bitów, co umożliwia efektywne przechowywanie i przesyłanie danych. 

## Zastosowanie
Oto kilka przykładów zastosowań w kontekście bezpieczeństwa base128:

- Ukrywanie danych:
   - Base128 może być używany do ukrywania danych, które nie są zrozumiałe dla zwykłego oka lub w przypadku przechwycenia przez osobę trzecią. Kodując dane za pomocą Base128, można zapobiec nieuprawnionemu dostępowi do wrażliwych informacji, co ma duże znaczenie w bezpieczeństwie danych.
- Utrudnianie odczytu i manipulacj:
   - Base128 utrudnia odczytanie i manipulację danych przez niepowołane osoby, ponieważ jest to mniej intuicyjne niż inne popularne formaty kodowania, takie jak Base64. Oznacza to, że jeśli dane zostaną przechwycone przez niepowołane osoby, trudniej będzie im zrozumieć i zmodyfikować te dane.
- Transport danych:
   - Może być używany do bezpiecznego transportu danych między różnymi składnikami systemu. Obejmuje to bezpieczne przesyłanie haseł, tokenów uwierzytelniających oraz innych wrażliwych informacji przez sieć, gdzie niepożądane odczytywanie lub zmienianie danych stanowi zagrożenie.
- Pseudonimizacja danych:
   -  Służy jako narzędzie do pseudonimizacji danych, pozwalając na dostęp do danych przedstawionych w innym, nieodwracalnym formacie. Można to wykorzystać w kontekście analizy danych i zabezpieczania wrażliwych informacji przed nieuprawnionym dostępem.


## Instalacja 
Aby zainstalować Base128, wykonaj następujące kroki: 

1. Sklonuj repozytorium z GitHuba: 

``` git clone https://github.com/username/base128.git ``` 

2. Przejdź do katalogu z projektem: ``` cd base128 ``` 

3. Skompiluj kod źródłowy: ``` make ``` 

4. Gotowe!
   
6. Base128 jest teraz gotowe do użycia. 

## Jak używać ### 
#### Kodowanie ####
Aby zakodować dane przy użyciu Base128, wykonaj polecenie: 

``` ./base128 -e ``` Gdzie: - ``` jest ścieżką do pliku, który chcesz zakodować, 

- `` jest ścieżką docelową, gdzie zostanie zapisany plik wyjściowy z zakodowanymi danymi. 

#### Dekodowanie ####
Aby zdekodować dane zakodowane w formacie Base128, wykonaj polecenie: 

``` ./base128 -d ``` Gdzie: - `` jest ścieżką do pliku, który chcesz zdekodować, 
- `` jest ścieżką docelową, gdzie zostanie zapisany plik wyjściowy z zdekodowanymi danymi. 

## Przykłady ### 
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

## Documentation 
See these PDF

## Download
Windows|Linux
-|-
ZIP|TAR.GZ
MD5|MD5


## License
[MIT License](/https://opensource.org/license/mit/)
