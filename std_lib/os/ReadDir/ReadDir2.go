package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var counts = make(map[string]int)

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

func main() {
	readDir("./", ".txt")
	for s, i := range counts {
		if i > 1 {
			fmt.Println(s, i)
		}
	}
}

// Read dir and return only filtered files contained within the folder
func readDir(dir string, ext string) {
	dirsAndFiles, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, dirOrFile := range dirsAndFiles {
		if !dirOrFile.IsDir() {
			if filepath.Ext(dirOrFile.Name()) == ".txt" {
				//fmt.Println(dirOrFile.Name(), " is a txt file!!!")
				f, err := os.Open(dirOrFile.Name())
				if err != nil {
					fmt.Fprintf(os.Stderr, "ReadDir: %v\n", err)
					continue
				}
				countLines(f, counts)
				f.Close()
			}
		}
	}
}