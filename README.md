# Base128 jest narzędziem służącym do kodowania i dekodowania danych w formacie Base128. #
Działa na zasadzie reprezentacji liczbowej danych wejściowych za pomocą 7-bitowych grup bitów, co umożliwia efektywne przechowywanie i przesyłanie danych. 

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

## License
MIT License
