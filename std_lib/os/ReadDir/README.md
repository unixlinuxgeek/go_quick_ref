package main

import (
"fmt"
"os"
"log"
)

func main() {
entries, err := os.ReadDir("./")
if err != nil {
log.Fatal(err)
}

    for _, e := range entries {
            fmt.Println(e.Name())
    }
}

