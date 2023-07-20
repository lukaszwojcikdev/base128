/**************************************************************************************************************************

W tym kodzie używamy paczki flag w celu analizy flagi --file, 
która pozwala na podanie ścieżki do pliku wejściowego przy uruchomieniu programu. 
Następnie otwieramy plik o podanej ścieżce i generujemy sumę kontrolną MD5 tak, 
jak wcześniej.

Aby użyć tego kodu, skopiuj go do pliku o rozszerzeniu ".go" (np. "md5sum.go") i uruchom wiersz poleceń. 
Następnie użyj polecenia "go run md5sum.go --file={ścieżka_do_pliku}" 
lub skompiluj program za pomocą polecenia "go build"” i uruchom skompilowaną binarkę z opcją "–file={ścieżka_do_pliku}". 

Przykład:

"go run md5sum.go --file=suma.txt" lub "./md5sum --file=suma.txt".

****************************************************************************************************************************/
package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// Deklaracja flagi dla ścieżki do pliku wejściowego
	filePath := flag.String("file", "", "sciezka do pliku wejsciowego")
	flag.Parse()

	// Sprawdzenie, czy ścieżka do pliku została podana
	if *filePath == "" {
		log.Fatal("Uzyj flagi --file , np. --file=plik.txt")
	}

	// Otwarcie pliku wejściowego
	file, err := os.Open(*filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Inicjalizacja obiektu MD5
	hash := md5.New()

	// Skopiowanie zawartości pliku do obiektu MD5
	if _, err := io.Copy(hash, file); err != nil {
		log.Fatal(err)
	}

	// Pobranie sumy kontrolnej MD5 jako bajtów
	hashBytes := hash.Sum(nil)

	// Konwersja bajtów do reprezentacji szesnastkowej
	hashString := fmt.Sprintf("%x", hashBytes)

	// Wyświetlenie sumy kontrolnej MD5
	fmt.Println("Suma kontrolna MD5:", hashString)
}
