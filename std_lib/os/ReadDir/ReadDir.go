// ReadDir читает указанный каталог, возвращая все его записи каталога, отсортированные по имени файла.
// Если при чтении каталога возникает ошибка, ReadDir возвращает записи,
// которые он мог прочитать до ошибки, вместе с самой ошибкой.

package main

import(
	"fmt"
	"os"
	"log"
)

func main() {
	readDir("./")
}

func readDir(dir string) {
	dirs, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, d := range dirs {
		fmt.Println(d)
	}
}